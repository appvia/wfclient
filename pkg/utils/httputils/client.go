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

package httputils

import (
	"net"
	"net/http"
	"os"
	"time"
)

var (
	DefaultTransport  *http.Transport
	DefaultHTTPClient *http.Client
)

func init() {
	DefaultTransport = http.DefaultTransport.(*http.Transport).Clone()
	DefaultTransport.DialContext = (&net.Dialer{
		Timeout: 5 * time.Second,
	}).DialContext
	DefaultTransport.TLSHandshakeTimeout = 5 * time.Second
	DefaultHTTPClient = NewDefaultHTTPClient(DefaultTransport)
}

// NewDefaultHTTPClient returns an initialised HTTP client using the provided transport. If the
// environment variable WAYFINDER_HTTP_CLIENT_TIMEOUT is set, that will be parsed as a timeout,
// otherwise the default value of 30 seconds will be used.
func NewDefaultHTTPClient(transport http.RoundTripper) *http.Client {
	if transport == nil {
		transport = DefaultTransport
	}
	c := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}
	if os.Getenv("WAYFINDER_HTTP_CLIENT_TIMEOUT") == "" {
		return c
	}
	timeout, err := time.ParseDuration(os.Getenv("WAYFINDER_HTTP_CLIENT_TIMEOUT"))
	if err == nil {
		c.Timeout = timeout
	}

	return c
}
