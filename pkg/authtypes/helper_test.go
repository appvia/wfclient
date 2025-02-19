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
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"

	jwtutils "github.com/appvia/wfclient/pkg/utils/jwt"
)

func TestIsExchangeScoped(t *testing.T) {
	cases := []struct {
		Claims   *jwtutils.Claims
		Expected bool
	}{
		{
			Claims:   jwtutils.NewClaims(jwt.MapClaims{}),
			Expected: false,
		},
		{
			Claims:   jwtutils.NewClaims(jwt.MapClaims{"aud": Audience}),
			Expected: false,
		},
		{
			Claims: jwtutils.NewClaims(jwt.MapClaims{
				"aud":    Audience,
				"scopes": []string{},
			}),
			Expected: false,
		},
		{
			Claims: jwtutils.NewClaims(jwt.MapClaims{
				"aud":    Audience,
				"scopes": []string{"me"},
			}),
			Expected: false,
		},
		{
			Claims: jwtutils.NewClaims(jwt.MapClaims{
				"aud":    Audience,
				"scopes": []string{ScopeExchange},
			}),
			Expected: true,
		},
	}

	for _, c := range cases {
		v := IsExchangeScoped(c.Claims)
		assert.Equal(t, c.Expected, v)
	}
}
