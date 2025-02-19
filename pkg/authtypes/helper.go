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

import (
	"github.com/appvia/wfclient/pkg/utils"
	jwsutils "github.com/appvia/wfclient/pkg/utils/jwt"
)

// IsTokenExpired checks the token for expiration
func IsTokenExpired(t string) (bool, error) {
	if t == "" {
		return true, nil
	}

	token, err := jwsutils.NewClaimsFromRawToken(t)
	if err != nil {
		return false, err
	}

	return token.HasExpired(), nil
}

// IsExchangeToken checks if the token is for exchange
func IsExchangeToken(token []byte) (bool, error) {
	claims, err := jwsutils.NewClaimsFromRawBytesToken(token)
	if err != nil {
		return false, err
	}

	return IsExchangeScoped(claims), nil
}

// IsAccessToken checks if the token is an access token
func IsAccessToken(token []byte) (bool, error) {
	claims, err := jwsutils.NewClaimsFromRawBytesToken(token)
	if err != nil {
		return false, err
	}

	return IsAccessTokenScoped(claims), nil
}

// IsExchangeScoped checks if the has the scopes for a platform or workspace access token exchange
func IsExchangeScoped(claims *jwsutils.Claims) bool {
	scopes, found := claims.GetScopes()
	if !found {
		return false
	}

	return utils.Contains(ScopeExchange, scopes)
}

// IsAccessTokenScoped checks if the has the scopes for a platform or workspace access token
func IsAccessTokenScoped(claims *jwsutils.Claims) bool {
	scopes, found := claims.GetScopes()
	if !found {
		return false
	}

	return utils.Contains(ScopeAccessToken, scopes)
}
