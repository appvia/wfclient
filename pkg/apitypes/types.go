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

package types

const (
	// APIBaseVersion is the current version for our non-resource API (i.e. our non-CRD API endpoints such as login)
	APIBaseVersion = "v2"

	// APIBasePath is the base path for the non-resource API (i.e. our non-CRD API endpoints such as login)
	APIBasePath = "/api/" + APIBaseVersion

	// ResourceAPIBasePath is the base path for the resource API (i.e. access to our CRDs)
	ResourceAPIBasePath = "/resources"

	// KubeProxyAPIBasePath is the base path for the kube proxy API (i.e. access to managed clusters)
	KubeProxyAPIBasePath = "/kubeproxy"

	// KubernetesResourceGroup is the group used to represent built-in Kubernetes resources in our
	// API, such as ConfigMaps and Secrets
	KubernetesResourceGroup = "k8s.io"
)

type LoginResponse struct {
	// IssuedToken
	IssuedToken *IssuedToken
}

// IssuedToken is a minted token from wayfinder
type IssuedToken struct {
	// RefreshToken is a token to retrieve new tokens, populated only on
	// initial login
	RefreshToken string `json:"refresh-token,omitempty"`
	// Token is the actual token for accessing the API
	Token string `json:"token,omitempty"`
	// Expires is the time token will expire
	Expires int64 `json:"expires,omitempty"`
}

// LocalUser represents a local user for login purposes
type LocalUser struct {
	// Username is the user's username
	Username string `json:"username,omitempty"`
	// Password used to identify the local user
	Password string `json:"password,omitempty"`
}

// WhoAmI provides a description to who you are
type WhoAmI struct {
	// AuthMethod is the authentication method being used
	AuthMethod string `json:"authMethod,omitempty"`
	// Username is your username
	Username string `json:"username,omitempty"`
	// Workspaces is a collection of workspaces you're a member of
	Workspaces []string `json:"workspaces,omitempty"`
	// WayfinderGroups is the list of Wayfinder-scoped groups that you are a member of
	WayfinderGroups []string `json:"wayfinderGroups,omitempty"`
	// WorkspaceGroups is a map of a workspace key to the list of groups you are a member of in that
	// workspace
	WorkspaceGroups map[string][]string `json:"workspaceGroups,omitempty"`
	// SubjectKind lets you know what you ARE
	SubjectKind string `json:"subjectKind,omitempty"`
}

// ServerInfo provides details around the settings and version of api
type ServerInfo struct {
	// FeatureFlags is a collection of supported / enabled features
	FeatureFlags map[string]bool `json:"featureFlags,omitempty"`
	// Version is the api version details
	Version Version `json:"version"`
	// InstanceIdentifier is the unique identifier of Wayfinder
	InstanceIdentifier string `json:"instanceIdentifier"`
	// Namespace is the namespace Wayfinder is running in on the host management cluster
	Namespace string `json:"namespace"`
}

// Version defines the version of the api
type Version struct {
	// Release is the release tag
	Release string `json:"release"`
	// SHA is the git sha
	SHA string `json:"sha"`
}
