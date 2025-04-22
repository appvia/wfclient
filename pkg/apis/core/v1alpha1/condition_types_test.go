/**
 * Copyright 2021 Appvia Ltd <info@appvia.io>
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

package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConditionsAreSame(t *testing.T) {
	tests := []struct {
		name       string
		a          Conditions
		b          Conditions
		expectSame bool
	}{
		{name: "Empty should be equal", expectSame: true},
		{
			name: "Same conditions with same status should be equal",
			a: Conditions{
				{Type: "bob", Status: "True"},
				{Type: "simon", Status: "True"},
				{Type: "jennifer", Status: "False"},
			},
			b: Conditions{
				{Type: "bob", Status: "True"},
				{Type: "simon", Status: "True"},
				{Type: "jennifer", Status: "False"},
			},
			expectSame: true,
		},
		{
			name: "Same conditions with same status should be equal even if messages / details differ",
			a: Conditions{
				{Type: "bob", Status: "True", Message: "yo"},
				{Type: "simon", Status: "True", Detail: "deets"},
				{Type: "jennifer", Status: "False"},
			},
			b: Conditions{
				{Type: "bob", Status: "True"},
				{Type: "simon", Status: "True", Message: "deets"},
				{Type: "jennifer", Status: "False", Detail: "yo"},
			},
			expectSame: true,
		},
		{
			name: "Should be unequal if status of a condition differs",
			a: Conditions{
				{Type: "bob", Status: "True", Message: "yo"},
				{Type: "simon", Status: "True", Detail: "deets"},
				{Type: "jennifer", Status: "False"},
			},
			b: Conditions{
				{Type: "simon", Status: "False", Message: "deets"},
				{Type: "bob", Status: "True"},
				{Type: "jennifer", Status: "False", Detail: "yo"},
			},
			expectSame: false,
		},
		{
			name: "Should be unequal if condition in A but not B",
			a: Conditions{
				{Type: "bob", Status: "True", Message: "yo"},
				{Type: "simon", Status: "True", Detail: "deets"},
				{Type: "jennifer", Status: "False"},
			},
			b: Conditions{
				{Type: "bob", Status: "True"},
				{Type: "jennifer", Status: "False", Detail: "yo"},
			},
			expectSame: false,
		},
		{
			name: "Should be unequal if condition in B but not A",
			a: Conditions{
				{Type: "bob", Status: "True", Message: "yo"},
				{Type: "jennifer", Status: "False"},
			},
			b: Conditions{
				{Type: "bob", Status: "True"},
				{Type: "simon", Status: "True", Detail: "deets"},
				{Type: "jennifer", Status: "False", Detail: "yo"},
			},
			expectSame: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectSame, tt.a.AreSame(tt.b))
		})
	}
}
