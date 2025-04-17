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

package config

import (
	"path"

	osutils "github.com/appvia/wfclient/pkg/utils/os"
)

var (
	// DefaultWayfinderConfigPath is the default path for the wayfinder configuration file
	DefaultWayfinderConfigPath = path.Join(osutils.UserHomeDir(), ".wayfinder", "config")
	// DefaultWayfinderConfigPathEnv is the default name of the env variable for config
	DefaultWayfinderConfigPathEnv = "WAYFINDER_CONFIG"
)

// Config is the configuration for the api
type Config struct {
	// AuthInfos is a collection of credentials
	AuthInfos map[string]*AuthInfo `json:"users,omitempty" yaml:"users,omitempty"`
	// CurrentProfile is the profile in use at the moment
	CurrentProfile string `json:"current-profile,omitempty" yaml:"current-profile,omitempty"`
	// Profiles is a collection of profiles
	Profiles map[string]*Profile `json:"profiles,omitempty" yaml:"profiles,omitempty"`
	// Servers is a collection of api endpoints
	Servers map[string]*Server `json:"servers,omitempty" yaml:"servers,omitempty"`
	// Version is the version of the configuration
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

// AuthInfo defines a credential to the api endpoint
type AuthInfo struct {
	// Identity is a wayfinder managed identity
	Identity *Identity `json:"identity,omitempty" yaml:"identity,omitempty"`
	// Token is a static token to use
	Token *string `json:"token,omitempty" yaml:"token,omitempty"`
}

// Identity is a wayfinder manage identity
type Identity struct {
	// RefreshToken represents a wayfinder managed refresh token issued by the wayfinder
	RefreshToken string `json:"refresh-token,omitempty" yaml:"refresh-token,omitempty"`
	// Token represents a wayfinder managed token issued by the wayfinder service
	Token string `json:"token,omitempty" yaml:"token,omitempty"`
}

// Profile links endpoint and a credential together
type Profile struct {
	// AuthInfo is the credentials to use
	AuthInfo string `json:"user,omitempty" yaml:"user,omitempty"`
	// Server is a reference to the server config
	Server string `json:"server,omitempty" yaml:"server,omitempty"`
	// Workspace is the default workspace for this profile
	Workspace string `json:"workspace,omitempty" yaml:"workspace,omitempty"`
}

// Server defines an endpoint for the api server
type Server struct {
	// Endpoint the url for the api endpoint of wayfinder
	Endpoint string `json:"server,omitempty" yaml:"server,omitempty"`
	// CACertificate is the ca bundle used to verify a self-signed api
	CACertificate string `json:"caCertificate,omitempty" yaml:"caCertificate,omitempty"`
	// APIInfo is a set of metadata about this instance of Wayfinder
	APIInfo *APIInfo `json:"apiInfo,omitempty"`
}

const defaultAPIBase = "/api/v2"
const defaultResourceBase = "/resources"
const defaultKubeProxyBase = "/kubeproxy"

func (s *Server) GetAPIInfo() APIInfo {
	if s.APIInfo != nil {
		return *s.APIInfo
	}
	// Where we have no API info, return something hard-coded to work with v2.4+ WF
	return APIInfo{
		NonResourceAPI: defaultAPIBase,
		ResourceAPI:    defaultResourceBase,
		KubeProxyAPI:   defaultKubeProxyBase,
	}
}

// APIInfo is a representation of the structure returned from the API server on the unauthenticated
// /apiinfo endpoint. This needs to be kept in sync with the APIInfo struct from
// /pkg/apiserver/types/types.go
type APIInfo struct {
	// NonResourceAPI is the base path for the non-resource API (i.e. our non-CRD API endpoints such
	// as login)
	NonResourceAPI string `json:"nonResourceAPI,omitempty"`
	// ResourceAPI is the base path for the resource API (i.e. access to our CRDs)
	ResourceAPI string `json:"resourceAPI,omitempty"`
	// KubeProxyAPI is the base path for the kube proxy API (i.e. access to managed clusters)
	KubeProxyAPI string `json:"kubeProxyAPI,omitempty"`
}
