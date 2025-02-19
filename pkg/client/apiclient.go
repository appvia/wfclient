/**
 * Copyright 2025 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	corev1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
	"github.com/appvia/wfclient/pkg/client/config"
	"github.com/appvia/wfclient/pkg/common"
	"github.com/appvia/wfclient/pkg/utils/httputils"
	"github.com/appvia/wfclient/pkg/utils/retry"
	"github.com/appvia/wfclient/pkg/utils/validation"
	"github.com/appvia/wfclient/pkg/version"
)

// apiClient implements the api and raw client
type apiClient struct {
	// body is the response body
	body *bytes.Buffer
	// ctx is the context for the client
	ctx context.Context
	// cfg is the client configuration
	cfg *config.Config
	// authtoken is a authorization header
	authtoken string
	// ferror is used to handle errors in the method chain
	ferror error
	// handler is an update handler for the config
	handler UpdateHandlerFunc
	// hc is the http client to use
	hc *http.Client
	// profile is the name of the profile to use
	profile string
	// payload is the outbound payload
	payload interface{}
	// response is the raw http response
	response *http.Response
	// result is what we decode into
	result interface{}
	// follow indiates we convert to websocket and follow stream
	follow bool
	// client is a reference to the main client interface
	client Interface
	// unauthenticated indicates no need to add auth
	unauthenticated bool

	urlManager URLManager
	// warningHandler defines how the client deals with response warnings
	warningHandler WarningHandler

	// customRequestDo is a function to perform the request, exposed so we can override it when testing.
	customRequestDo RequestDo
}

func (a *apiClient) Profile() string {
	return a.profile
}

// makeAPIEndpoint is responsible for getting the api endpoint
func (a *apiClient) makeAPIEndpoint() (*config.Server, error) {
	// @step: check we have the endpoint
	profile, found := a.cfg.Profiles[a.Profile()]
	if !found {
		return nil, ErrMissingProfile
	}
	server, found := a.cfg.Servers[profile.Server]
	if !found {
		return nil, NewProfileInvalidError("missing profile server", a.Profile())
	}

	if server.Endpoint == "" {
		return nil, NewProfileInvalidError("missing endpoint", a.Profile())

	}

	return server, nil
}

func (a *apiClient) reqCtx() context.Context {
	if a.ctx != nil {
		return a.ctx
	}
	return context.Background()
}

// handleRequest is responsible for handling the request chain
func (a *apiClient) handleRequest(method string) RestInterface {
	ctx := a.reqCtx()
	err := func() error {
		if a.ferror != nil {
			return a.ferror
		}

		// @step: find the endpoint
		server, err := a.makeAPIEndpoint()
		if err != nil {
			a.ferror = err

			return a.ferror
		}

		if a.hc == nil && a.customRequestDo == nil {
			a.hc = a.makeHTTPClient(server.CACertificate)
			if a.follow {
				a.hc.Timeout = 0 * time.Second
			}
		}

		// @step: we generate the uri from the parameter
		uri, err := a.urlManager.MakeURL(server.GetAPIInfo())
		if err != nil {
			a.ferror = err

			return a.ferror
		}
		logFields := map[string]interface{}{
			"endpoint": server.Endpoint,
			"method":   method,
			"uri":      uri,
		}
		if server.CACertificate != "" {
			logFields["customCA"] = true
		}

		common.Log(ctx).WithFields(logFields).Debug("API request")

		// @step: we generate the fully qualified url
		ep := fmt.Sprintf("%s/%s", server.Endpoint, uri)

		now := time.Now()

		var resp *http.Response

		err = retry.Retry(ctx, 3, true, 1*time.Second, func() (bool, error) {
			var respErr error
			resp, respErr = a.makeRequest(method, ep)
			if respErr != nil {
				return false, respErr
			}

			if resp.StatusCode == http.StatusTooManyRequests {
				common.Log(ctx).WithFields(logFields).
					Warn("API request: Received 'Too many requests' error, backing off and retrying")
				return false, nil
			}

			return true, nil
		})

		if err != nil && err != retry.ErrReachMaxAttempts {
			common.Log(ctx).WithFields(logFields).WithError(err).WithField("duration", time.Since(now).String()).Debug("API request: Error")
			return err
		}

		common.Log(ctx).WithFields(logFields).WithField("reponseCode", resp.StatusCode).WithField("duration", time.Since(now).String()).Debug("API request: Complete")

		return a.handleResponse(resp)
	}()
	if err != nil {
		a.ferror = err
	}

	return a
}

// makeRequest is responsible for preparing and handling the http request
func (a *apiClient) makeRequest(method, url string) (*http.Response, error) {
	// @step: do we have any thing to encode?
	payload, err := a.makePayload()
	if err != nil {
		return nil, err
	}

	// @step: construct the http request
	ctx := a.reqCtx()
	request, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(ClientVersionHeader, version.Release)

	// @step: add the authentication from profile
	if err := a.AddAuthorization(request); err != nil {
		return nil, err
	}

	return a.do(request)
}

func (a *apiClient) do(req *http.Request) (*http.Response, error) {
	if a.customRequestDo != nil {
		return a.customRequestDo(req)
	}
	return a.hc.Do(req)
}

// Authorization allows you to override the authorization
func (a *apiClient) Authorization(h string) RestInterface {
	a.authtoken = h

	return a
}

// Unauthenticated indicates no need to add authorization
func (a *apiClient) Unauthenticated() RestInterface {
	a.unauthenticated = true

	return a
}

// AddAuthorization is responsible for adding the authorization headers
func (a *apiClient) AddAuthorization(req *http.Request) error {
	if a.unauthenticated {
		return nil
	}

	if a.authtoken != "" {
		req.Header.Set("Authorization", "Bearer "+a.authtoken)

		return nil
	}

	auth := a.cfg.AuthInfos[a.Profile()]

	switch {
	case auth == nil:
		return NewProfileInvalidError("missing authentication profile", a.Profile())

	case auth.Token != nil:
		req.Header.Set("Authorization", "Bearer "+*auth.Token)

	case auth.Identity != nil:

		expired, err := auth.Identity.IsExpired()
		if err != nil {
			return err
		}
		if expired {
			if err := a.client.RefreshIdentity(); err != nil {
				return err
			}
		}

		req.Header.Set("Authorization", "Bearer "+auth.Identity.Token)
	}

	return nil
}

// handleResponse is responsible for handling the http response from api
func (a *apiClient) handleResponse(resp *http.Response) error {
	// @step: if everything is ok, check for a response and return
	code := resp.StatusCode
	a.response = resp
	if code >= http.StatusOK && code <= 299 {
		if !a.follow {
			if err := a.makeResult(resp, a.result); err != nil {
				return err
			}
		}

		return nil
	}

	// @step: we have encountered an error we need read in a APIError or create one
	apiError := &APIError{}
	apiError.Code = resp.StatusCode
	apiError.Verb = resp.Request.Method
	apiError.URI = resp.Request.RequestURI

	a.decodeError(resp, apiError)

	if apiError.Message == "" {
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			apiError.Message = "Authorization required"
		case http.StatusNotFound:
			apiError.Message = "Resource does not exist"
		case http.StatusForbidden:
			apiError.Message = "Request denied, check your permissions"
		case http.StatusBadRequest:
			apiError.Message = "Invalid request"
		case http.StatusTooManyRequests:
			apiError.Message = "Too many requests, please try again shortly"
		case http.StatusServiceUnavailable:
			apiError.Message = "API service unavailable"
		case http.StatusMethodNotAllowed:
			apiError.Message = "Resource does not support method " + apiError.Verb
		case http.StatusConflict:
			apiError.Message = ObjectModifiedError
		default:
			apiError.Message = "Unexpected error from API"
		}
	}

	return apiError
}

func (a *apiClient) decodeError(resp *http.Response, apiError *APIError) {
	if resp.Body != nil {
		switch resp.StatusCode {
		case http.StatusBadRequest:
			vError := &validation.Error{}
			if err := a.makeResult(resp, vError); err != nil {
				common.Log(a.reqCtx()).WithError(err).Debug("response cannot be decoded into a validation error")
				return
			}
			apiError.Message = vError.Error()
			apiError.Validation = vError
			return
		case http.StatusConflict:
			// Two different types of conflict are represented by 409 - a conflict when trying to
			// write an object to k8s and a conflict with a dependency blocking deletion
			if resp.Header.Get("x-wayfinder-objectmodified") == "true" {
				apiError.Message = ObjectModifiedError
			} else {
				err := &validation.ErrDependencyViolation{}
				if err := a.makeResult(resp, err); err != nil {
					common.Log(a.reqCtx()).WithError(err).Debugf("response cannot be decoded into a validation error - %v", a.Body())
					return
				}
				apiError.Message = err.Error()
				apiError.DependencyViolation = err
			}
			return
		}

		err := a.makeResult(resp, apiError)
		if err != nil && err != io.EOF {
			common.Log(a.reqCtx()).WithError(err).Debug("response cannot be decoded")
		}
	}
}

// makeResult is responsible for reading the resulting payload
func (a *apiClient) makeResult(resp *http.Response, data interface{}) error {
	a.body = &bytes.Buffer{}

	if resp.Body == nil {
		common.Log(a.reqCtx()).Trace("request to api had no response body present")

		return nil
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	a.body.Write(content)

	common.Log(a.reqCtx()).WithField("body", a.body.String()).Trace("we received the following response from api")

	if data == nil || len(content) == 0 {
		common.Log(a.reqCtx()).Trace("no result has been set to save the payload")

		return nil
	}

	return json.NewDecoder(a.Body()).Decode(data)
}

// makePayload is responsible for encoding the payload if any
func (a *apiClient) makePayload() (io.Reader, error) {
	if a.payload == nil {
		return nil, nil
	}
	b := &bytes.Buffer{}

	if err := json.NewEncoder(b).Encode(a.payload); err != nil {
		return nil, err
	}
	common.Log(a.reqCtx()).WithField("payload", b.String()).Trace("using the attached payload for request")

	return b, nil
}

// Exists check is the resource exists
func (a *apiClient) Exists() (bool, error) {
	if err := a.Get().Error(); err != nil {
		if !IsNotFound(err) {
			return false, err
		}

		return false, nil
	}

	return true, nil
}

// Delete performs a delete
func (a *apiClient) Delete() RestInterface {
	request := a.handleRequest(http.MethodDelete)

	a.handleWarnings()

	return request
}

// Get performs a get request
func (a *apiClient) Get() RestInterface {
	return a.handleRequest(http.MethodGet)
}

// Create performs a get request
func (a *apiClient) Create() RestInterface {
	request := a.handleRequest(http.MethodPost)

	a.handleWarnings()

	return request
}

// Update performs an put request
func (a *apiClient) Update() RestInterface {
	request := a.handleRequest(http.MethodPut)

	a.handleWarnings()

	return request
}

func (a *apiClient) handleWarnings() {
	warnings := a.GetWarnings()
	if len(warnings) == 0 {
		return
	}
	ctx := a.reqCtx()
	common.Log(ctx).WithField("warnings", warnings).Debug("API request: Warnings received")
	if a.warningHandler != nil {
		a.warningHandler(ctx, warnings)
	}
}

// Post performs a post request
func (a *apiClient) Post() RestInterface {
	return a.handleRequest(http.MethodPost)
}

// Parameters defines a list of parameters for the request
func (a *apiClient) Parameters(params ...ParameterFunc) RestInterface {
	if err := a.urlManager.Parameters(params...); err != nil {
		a.ferror = err
	}

	return a
}

// makeHTTPClient is responsible for creating the http client
func (a *apiClient) makeHTTPClient(cacert string) *http.Client {
	if cacert == "" {
		return httputils.DefaultHTTPClient
	}

	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	if ok := rootCAs.AppendCertsFromPEM([]byte(cacert)); !ok {
		common.Log(a.reqCtx()).Debug("no certs appended, using system certs only")
	}

	custTransport := httputils.DefaultTransport.Clone()
	if custTransport.TLSClientConfig == nil {
		custTransport.TLSClientConfig = &tls.Config{}
	}
	custTransport.TLSClientConfig.RootCAs = rootCAs
	return httputils.NewDefaultHTTPClient(custTransport)
}

func (a *apiClient) HasParameter(key string) (string, bool) {
	return a.urlManager.HasParameter(key)
}

func (a *apiClient) Resource(src VersionedResourceSource) RestInterface {
	a.urlManager.Resource(src)
	return a
}

func (a *apiClient) ResourceAPIVersion(v string) RestInterface {
	a.urlManager.ResourceAPIVersion(v)
	return a
}

func (a *apiClient) ResourceVersion(v string) RestInterface {
	a.urlManager.ResourceVersion(v)
	return a
}

// Workspace set the workspace
func (a *apiClient) Workspace(v corev1.WorkspaceKey) RestInterface {
	a.urlManager.Workspace(v)
	return a
}

// Name sets the resource name
func (a *apiClient) Name(v string) RestInterface {
	a.urlManager.Name(v)
	return a
}

// SubResource adds a subresource to the operation
func (a *apiClient) SubResource(v string) RestInterface {
	a.urlManager.SubResource(v)
	return a
}

// SubResourceName adds a subresource to the operation
func (a *apiClient) SubResourceName(v string) RestInterface {
	a.urlManager.SubResourceName(v)
	return a
}

// Endpoint defines the endpoint to use
func (a *apiClient) Endpoint(v string) RestInterface {
	a.urlManager.Endpoint(v)
	return a
}

// RawEndpoint defines the endpoint to use without any prefixing
func (a *apiClient) RawEndpoint(v string) RestInterface {
	a.urlManager.RawEndpoint(v)
	return a
}

// Payload set the payload of the request
func (a *apiClient) Payload(v interface{}) RestInterface {
	a.payload = v

	return a
}

// Result set the object which we should decode into
func (a *apiClient) Result(v interface{}) RestInterface {
	a.result = v

	return a
}

// Context sets the request context
func (a *apiClient) Context(ctx context.Context) RestInterface {
	a.ctx = ctx

	return a
}

// Error return any error and resets post
func (a *apiClient) Error() error {
	// we need to reset the error
	defer func() {
		a.ferror = nil
	}()

	return a.ferror
}

// Duplicate duplicates the current request
func (a *apiClient) Duplicate() RestInterface {
	n := &apiClient{
		cfg:             a.cfg,
		payload:         a.payload,
		profile:         a.profile,
		result:          a.result,
		client:          a.client,
		urlManager:      a.urlManager.Duplicate(),
		warningHandler:  a.warningHandler,
		customRequestDo: a.customRequestDo,
	}

	return n
}

// Follow indicates we follow the stream
func (a *apiClient) Follow(v bool) RestInterface {
	a.follow = v

	return a
}

// Do returns both the response and error
func (a *apiClient) Do() (RestInterface, error) {
	return a, a.ferror
}

// GetPayload returns the payload for inspection
func (a *apiClient) GetPayload() interface{} {
	return a.payload
}

// Body returns the body if any
func (a *apiClient) Body() io.Reader {
	if a.follow {
		return a.response.Body
	}
	return strings.NewReader(a.body.String())
}

func (a *apiClient) GetWarnings() []validation.Warning {
	warnings := []validation.Warning{}
	if a.response == nil || a.response.Header == nil {
		return warnings
	}
	warningStrings := a.response.Header.Values(validation.WarningHeader)

	for _, w := range warningStrings {
		warning := &validation.Warning{}

		if err := json.Unmarshal([]byte(w), warning); err != nil {
			common.Log(a.reqCtx()).WithError(err).Debug(fmt.Errorf("response warning can't be parsed into a validation warning: %s", w))
		} else {
			warnings = append(warnings, *warning)
		}
	}

	return warnings
}

func (a *apiClient) WithWarningHandler(handler WarningHandler) RestInterface {
	a.warningHandler = handler
	return a
}
