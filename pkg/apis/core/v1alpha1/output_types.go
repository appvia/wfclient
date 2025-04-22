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

type OutputDefinition struct {
	// Name is the name of the output
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Description is the description of the output
	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`

	// Sensitive indicates if the output is sensitive
	// +kubebuilder:validation:Optional
	Sensitive bool `json:"sensitive,omitempty"`
}

type Output struct {
	Name string `json:"name"`
	// Sensitive will be set if the plan identifies this output as a sensitive value. In this case,
	// Value will not be populated.
	Sensitive bool `json:"sensitive,omitempty"`
	// Value is the output value. This will not be populated for sensitive outputs.
	Value string `json:"value,omitempty"`
	// OutputType is the type of the output value. This is a hint to the consumer of the output as
	// to the format of the Value.
	OutputType string `json:"outputType,omitempty"`
}

// OutputAware can be implemented by Wayfinder resources that generate outputs so they can be read
// generically.
// +kubebuilder:object:generate=false
type OutputAware interface {
	// GetOutputs returns the available outputs for this resource. The values of any outputs marked
	// sensitive will not be included.
	GetOutputs() []Output
	// GetOutputsSecret returns the name of a secret (in the workspace if a workspaced resource,
	// else a platform secret for a non-workspaced resource) which contains all of the outputs
	// including sensitive values.
	GetOutputsSecret() string
}
