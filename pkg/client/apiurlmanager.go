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
	"fmt"
	"net/url"
	"strings"

	corev1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
	"github.com/appvia/wfclient/pkg/client/config"
)

// URLManager provides a consistent place for us to generate URLs for our API clients. This was
// predominantly separated out so we could share it between the fake API client and the real one,
// as this logic is equally applicable to both. It's also rather easier to unit test in isolation...
type URLManager struct {
	// parameters hold the parameters for the request
	parameters map[string]string
	// queryparams are a collection of query parameters
	queryparams url.Values
	// endpoint is the raw endpoint template to use
	endpoint string
	// rawEndpoint uses the provided endpoint without prefixing it with the server base path
	rawEndpoint bool
	// versionedResource changes the semantics to use versioned APIs to access the resource
	versionedResource bool
	// isVersionedCheck returns true if the provided API version is versioned
	isVersionedCheck func(ver string) bool
}

func NewURLManager() URLManager {
	return URLManager{
		parameters: make(map[string]string),
	}
}

const (
	paramWorkspace       = "workspace"
	paramAPIVersion      = "apiVersion"
	paramGroup           = "group"
	paramResource        = "resource"
	paramName            = "name"
	paramResourceVersion = "resourceVersion"
	paramSubresource     = "subresource"
	paramSubresourceName = "subresourcename"
)

func (a URLManager) Duplicate() URLManager {
	n := URLManager{
		queryparams:       a.queryparams,
		endpoint:          a.endpoint,
		rawEndpoint:       a.rawEndpoint,
		parameters:        make(map[string]string),
		versionedResource: a.versionedResource,
	}
	for k, v := range a.parameters {
		n.parameters[k] = v
	}
	return n
}

func (a URLManager) MakeURL(apiInfo config.APIInfo) (string, error) {
	if a.endpoint != "" {
		// A custom endpoint has been provided, use that.
		return a.makeEndpointURL(apiInfo), nil
	}

	// By default, we use resource-oriented URLs
	return a.makeResourceURL(apiInfo)
}

// IsResourceRequest will return true if this is not an endpoint request
func (a URLManager) IsResourceRequest() bool {
	return a.endpoint == ""
}

// IsSubResourceRequest will return true if this a request for a subresource
func (a URLManager) IsSubResourceRequest() bool {
	_, srFound := a.HasParameter(paramSubresource)
	return srFound
}

// GroupVersionKind will return a string identifying the GVK of a resource request
func (a URLManager) GetGroupVersionKind() (string, string, string) {
	if !a.IsResourceRequest() {
		return "", "", ""
	}
	g, gFound := a.HasParameter(paramGroup)
	v, vFound := a.HasParameter(paramAPIVersion)
	k, kFound := a.HasParameter(paramResource)
	if !gFound || !vFound || !kFound {
		return "", "", ""
	}

	return g, v, k
}

func (a URLManager) GetName() string {
	return a.parameters[paramName]
}

func (a URLManager) GetWorkspace() corev1.WorkspaceKey {
	ws, wsFound := a.HasParameter(paramWorkspace)
	if !wsFound {
		return ""
	}
	return corev1.WorkspaceKey(ws)
}

// MakeResourceURL generates a URL in the format
// /api/<group>/<version>/workspaces/<workspace>/<kind>/<name>/<subresource>/<subresourcename> or
// /api/<group>/<version>/<kind>/<name>/<subresource>/<subresourcename> as appropriate for the request
func (a URLManager) makeResourceURL(apiInfo config.APIInfo) (string, error) {
	g, gFound := a.HasParameter(paramGroup)
	v, vFound := a.HasParameter(paramAPIVersion)
	if !gFound || !vFound {
		return "", fmt.Errorf("unable to determine API group and version for resource, cannot perform API operation")
	}
	paths := []string{strings.TrimPrefix(apiInfo.ResourceAPI, "/"), g, v}

	if value, found := a.HasParameter(paramWorkspace); found && value != "" {
		paths = append(paths, []string{"workspaces", value}...)
	}
	for _, x := range []string{paramResource, paramName} {
		if value, found := a.HasParameter(x); found {
			paths = append(paths, value)
		}
	}

	// If we have a versioned resource with a name, add either /versions or /versions/VERSION n to the path.
	if _, hasName := a.HasParameter(paramName); a.versionedResource && hasName {
		paths = append(paths, "versions")
		if rv, found := a.HasParameter(paramResourceVersion); found {
			paths = append(paths, rv)
		}
	}

	for _, x := range []string{paramSubresource, paramSubresourceName} {
		if value, found := a.HasParameter(x); found {
			paths = append(paths, value)
		}
	}

	// @step: we add the path elements and the queries together
	uri := strings.Join(paths, "/")
	if len(a.queryparams) > 0 {
		uri = fmt.Sprintf("%s?%s", uri, a.queryparams.Encode())
	}

	return uri, nil
}

// MakeEndpointURL is responsible for generating the url from a template
func (a URLManager) makeEndpointURL(apiInfo config.APIInfo) string {
	var uri string
	if !a.rawEndpoint {
		uri = apiInfo.NonResourceAPI + "/" + strings.TrimPrefix(a.endpoint, "/")
	} else {
		uri = a.endpoint
	}

	// Strip any leading /
	uri = strings.TrimPrefix(uri, "/")

	// @step: we add the path elements and the queries together
	for param, value := range a.parameters {
		uri = strings.ReplaceAll(uri, "{"+param+"}", value)
	}

	// @step: add the query params if any to the url
	if len(a.queryparams) > 0 {
		uri = fmt.Sprintf("%s?%s", uri, a.queryparams.Encode())
	}

	return uri
}

// SubResource adds a subresource to the operation
func (a *URLManager) SubResource(v string) {
	a.parameters[paramSubresource] = v
}

// SubResourceName adds a subresource to the operation
func (a *URLManager) SubResourceName(v string) {
	a.parameters[paramSubresourceName] = v
}

// Name sets the resource name
func (a *URLManager) Name(v string) {
	if v == "" {
		return
	}
	a.parameters[paramName] = v
}

func (a *URLManager) Resource(src VersionedResourceSource) {
	if src == nil {
		return
	}
	a.isVersionedCheck = src.IsResourceVersioned
	a.parameters[paramResource] = src.GetAPIName()
	gv := src.GetGroupVersion()
	if gv.Group != "" {
		a.parameters[paramGroup] = gv.Group
	}
	if gv.Version != "" {
		a.parameters[paramAPIVersion] = gv.Version
	}
	a.versionedResource = a.isVersionedCheck(gv.Version)
}

func (a *URLManager) ResourceAPIVersion(v string) {
	if v == "" {
		return
	}
	a.parameters[paramAPIVersion] = v
	if a.isVersionedCheck != nil {
		a.versionedResource = a.isVersionedCheck(v)
	}
}

func (a *URLManager) ResourceVersion(rv string) {
	if rv == "" {
		return
	}
	a.parameters[paramResourceVersion] = rv
}

func (a *URLManager) Workspace(v corev1.WorkspaceKey) {
	if v != "" {
		a.parameters[paramWorkspace] = string(v)
	}
}

// HasParameter checks if the parameter exists
func (a *URLManager) HasParameter(key string) (string, bool) {
	value, found := a.parameters[key]

	return value, (found && value != "")
}

// HasQueryParameter checks if the query parameter exists
func (a *URLManager) HasQueryParameter(key string) ([]string, bool) {
	value, found := a.queryparams[key]

	return value, (found && len(value) > 0)
}

// Parameters defines a list of parameters for the request
func (a *URLManager) Parameters(params ...ParameterFunc) error {
	for _, fn := range params {
		param, err := fn()
		if err != nil {
			return err
		}
		if param.IsPath {
			a.parameters[param.Name] = param.Value
		} else {
			if a.queryparams == nil {
				a.queryparams = url.Values{}
			}
			a.queryparams.Add(param.Name, param.Value)
		}
	}
	return nil
}

// Endpoint defines the endpoint to use
func (a *URLManager) Endpoint(v string) {
	a.endpoint = v
}

// RawEndpoint defines the endpoint to use without any prefixing
func (a *URLManager) RawEndpoint(v string) {
	a.endpoint = v
	a.rawEndpoint = true
}
