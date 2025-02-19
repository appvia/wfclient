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

package authtypes

// MemberRole is the name of the default member role
const MemberRole = "member"
const (
	// AuthenicationModeToken indicates a admin token
	AuthenicationModeToken = "token"
	// AuthenicationModeKubernetes indicates a internal service account
	AuthenicationModeKubernetes = "kubernetes"
	// AuthenicationModeSSO indicates an sso account
	AuthenicationModeSSO = "sso"
	// AuthenicationModeJWT indicates an jwt account
	AuthenicationModeJWT = "jwt"
	// SystemAuthenticated indicates the request has an authenticated user
	SystemAuthenticated = "system:authenticated"
	// WorkspaceAuthenticated indicates a member of request resource
	WorkspaceAuthenticated = "workspace:authenticated"
)

const (
	// Audience is the audience of the tokens
	Audience = "wayfinder"
	// KubernetesAudience is the audience of the tokens
	KubernetesAudience = "kubernetes"
	// RefreshTokenAudience is the audience for wayfinder refresh tokens
	RefreshTokenAudience = "refresh"
)

const (
	// ScopeExchange indicates a token exchange scope for an access token
	ScopeExchange = "wayfinder:auth:exchange"
	// ScopeRefresh indicates a refresh token
	ScopeRefresh = "wayfinder:auth:refresh"
	// ScopeCookieRefresh indicates a refresh token only valid for cookie-based exchanges
	ScopeCookieRefresh = "wayfinder:auth:cookierefresh"
	// ScopeAccessToken is a scope used for tokens for access tokens (both workspace and
	// wayfinder-scoped)
	ScopeAccessToken = "wayfinder:system:accesstoken"
	// ScopeKubernetesAccount indicates a kubernetes account
	ScopeKubernetesAccount = "wayfinder:system:kubernetes"
	// ScopeUser is a user account type
	ScopeUser = "wayfinder:system:user"
)
