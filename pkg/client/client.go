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
	"errors"
	"net/http"
	"time"

	"github.com/appvia/wfclient/pkg/client/config"
	"github.com/appvia/wfclient/pkg/common"
)

var (
	// ErrAuthenticationRequired requires authentication
	ErrAuthenticationRequired = APIError{
		Code:    http.StatusUnauthorized,
		Message: "authentication required",
	}
)

// cc provides a wrapper around th config
type cc struct {
	cfg            *config.Config
	handler        UpdateHandlerFunc
	warningHandler WarningHandler
	profile        string
	apiClient      func(cfg *config.Config) RestInterface
	requestDo      RequestDo
}

// NewClient returns a new client for the provided config, without silly nil checks for nicer usage.
// This will panic if you supply a nil config.
func NewClient(cfg *config.Config, options ...OptionFunc) Interface {
	c, err := New(cfg, options...)
	if err != nil {
		panic(err.Error())
	}
	return c
}

// New creates and returns an API client
func New(cfg *config.Config, options ...OptionFunc) (Interface, error) {
	if cfg == nil {
		return nil, errors.New("no client configuration")
	}

	c := &cc{cfg: cfg}

	// apply the options
	for _, fn := range options {
		fn(c)
	}
	if c.apiClient == nil {
		c.apiClient = func(cfg *config.Config) RestInterface {
			return &apiClient{
				cfg:             cfg,
				client:          c,
				handler:         c.handler,
				urlManager:      NewURLManager(),
				profile:         c.CurrentProfile(),
				warningHandler:  c.warningHandler,
				customRequestDo: c.requestDo,
			}
		}
	}

	return c, nil
}

// Config return a copy of the client configuration
func (c *cc) Config() *config.Config {
	return c.cfg
}

// OverrideProfile sets the default profile to use
func (c *cc) OverrideProfile(name string) Interface {
	c.profile = name

	return c
}

// CurrentProfile returns the current profile
func (c *cc) CurrentProfile() string {
	if c.profile != "" {
		return c.profile
	}

	return c.cfg.CurrentProfile
}

// Request creates a request instance
func (c *cc) Request() RestInterface {
	return c.apiClient(c.cfg)
}

// RefreshIdentity is called to refresh the identity token of the client
func (c *cc) RefreshIdentity() error {

	auth := c.cfg.AuthInfos[c.CurrentProfile()]

	switch {
	case auth == nil:
		return NewProfileInvalidError("missing authentication profile", c.CurrentProfile())
	case auth.Identity == nil:
		return errors.New("no token available to refresh")

	case auth.Identity.IsExchangeToken():
		common.LogWithoutContext().Debug("Refresh access token via access token exchange")

		ttl := 30 * time.Minute
		token, err := ExchangeAccessToken(c, []byte(auth.Identity.RefreshToken), ttl)
		if err != nil {
			common.LogWithoutContext().WithError(err).Error("trying to exchange access token")

			return err
		}
		auth.Identity.Token = string(token)

	case auth.Identity.RefreshToken != "":
		common.LogWithoutContext().Debug("Refresh identity token")

		token, err := RefreshWayfinderIdentityToken(c, []byte(auth.Identity.RefreshToken))
		if err != nil {
			return err
		}
		auth.Identity.Token = string(token)

	default:
		return errors.New("no refresh or exchange token available to refresh")
	}

	return c.handleConfigurationUpdate()
}

func (c *cc) CheckServer(force, saveProfile bool) error {
	prof := c.cfg.Profiles[c.CurrentProfile()]
	if prof == nil {
		return ErrMissingProfile
	}
	serv := c.cfg.Servers[prof.Server]
	if serv == nil {
		return ErrMissingProfile
	}
	// If already set, nothing to do.
	if !force && serv.APIInfo != nil {
		return nil
	}

	// Ping the API to check the versioning
	apiInfo := &config.APIInfo{}
	if err := c.Request().RawEndpoint("/apiinfo").Unauthenticated().Result(apiInfo).Get().Error(); err != nil {
		if IsNotFound(err) {
			// Backwards-compat for pre-2.4:
			apiInfo.NonResourceAPI = "/api/v1alpha1"
		} else {
			return err
		}
	}

	serv.APIInfo = apiInfo

	if !saveProfile {
		return nil
	}

	return c.handleConfigurationUpdate()
}

// handleConfigurationUpdate is called when the configuration has been updated
func (c *cc) handleConfigurationUpdate() error {
	if c.handler == nil {
		return nil
	}

	return c.handler()
}
