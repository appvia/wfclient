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

func TestInputDefinitions_Get(t *testing.T) {
	tests := []struct {
		name     string
		inputs   InputDefinitions
		findName string
		want     *InputDefinition
	}{
		{
			name: "finds existing input",
			inputs: InputDefinitions{
				{Name: "test1", Type: InputTypeString},
				{Name: "test2", Type: InputTypeNumber},
			},
			findName: "test1",
			want:     &InputDefinition{Name: "test1", Type: InputTypeString},
		},
		{
			name: "returns nil for non-existent input",
			inputs: InputDefinitions{
				{Name: "test1", Type: InputTypeString},
			},
			findName: "missing",
			want:     nil,
		},
		{
			name:     "returns nil for empty definitions",
			inputs:   InputDefinitions{},
			findName: "test",
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.inputs.Get(tt.findName)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInputValidationDefinition_IsRequired(t *testing.T) {
	trueVal := true
	falseVal := false

	tests := []struct {
		name       string
		validation InputValidationDefinition
		want       bool
	}{
		{
			name:       "returns false when Required is nil",
			validation: InputValidationDefinition{},
			want:       false,
		},
		{
			name:       "returns true when Required is true",
			validation: InputValidationDefinition{Required: &trueVal},
			want:       true,
		},
		{
			name:       "returns false when Required is false",
			validation: InputValidationDefinition{Required: &falseVal},
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.validation.IsRequired()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInputValues_Get(t *testing.T) {
	tests := []struct {
		name     string
		inputs   InputValues
		findName string
		want     *InputValue
	}{
		{
			name: "finds existing value",
			inputs: InputValues{
				{Name: "test1", Value: "value1"},
				{Name: "test2", Value: "42"},
			},
			findName: "test1",
			want:     &InputValue{Name: "test1", Value: "value1"},
		},
		{
			name: "returns nil for non-existent value",
			inputs: InputValues{
				{Name: "test1", Value: "value1"},
			},
			findName: "missing",
			want:     nil,
		},
		{
			name:     "returns nil for empty values",
			inputs:   InputValues{},
			findName: "test",
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.inputs.Get(tt.findName))
		})
	}
}

func TestInputValue_HasValue(t *testing.T) {
	tests := []struct {
		name  string
		input *InputValue
		want  bool
	}{
		{
			name:  "returns false for nil input",
			input: nil,
			want:  false,
		},
		{
			name:  "returns true for empty not nil value",
			input: &InputValue{Name: "test"},
			want:  true,
		},
		{
			name: "returns true for non-empty value",
			input: &InputValue{
				Name:  "test",
				Value: "value",
			},
			want: true,
		},
		{
			name: "returns true for non-empty value",
			input: &InputValue{
				Name:  "test",
				Value: "value",
			},
			want: true,
		},
		{
			name: "returns true for empty string value",
			input: &InputValue{
				Name:  "test",
				Value: "",
			},
			want: true,
		},
		{
			name: "returns true for zero integer value",
			input: &InputValue{
				Name:  "test",
				Value: "0",
			},
			want: true,
		},
		{
			name: "returns true for false boolean value",
			input: &InputValue{
				Name:  "test",
				Value: "false",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.HasValue()
			assert.Equal(t, tt.want, got)
		})
	}
}
