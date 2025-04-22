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

const (
	// InputTypeString is a string type
	InputTypeString = "string"
	// InputTypeListString is a list type of strings
	InputTypeListString = "list(string)"
	// InputTypeNumber is a number type
	InputTypeNumber = "number"
	// InputTypeBool is a boolean type
	InputTypeBool = "bool"
)

// InputTypes is a list of the possible input types
var InputTypes = []string{
	InputTypeString,
	InputTypeListString,
	InputTypeNumber,
	InputTypeBool,
}

// InputDefinition defines an input to a resource.
type InputDefinition struct {
	// Description is the description of the variable
	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`

	// Name is the name of the variable
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Sensitive indicates if the input is sensitive
	// +kubebuilder:validation:Optional
	Sensitive bool `json:"sensitive,omitempty"`

	// Type is the type of the variable
	// +kubebuilder:validation:Optional
	Type string `json:"type,omitempty"`

	// Validation is the validation of the variable
	// +kubebuilder:validation:Required
	Validation InputValidationDefinition `json:"validation"`

	// Options are the available options for a list or string type
	// +kubebuilder:validation:Optional
	Options []string `json:"options,omitempty"`
}

// InputDefinitions defines a list of input definitions
// +kubebuilder:validation:Type=array
type InputDefinitions []InputDefinition

func (ids InputDefinitions) Get(name string) *InputDefinition {
	for _, id := range ids {
		if id.Name == name {
			return &id
		}
	}

	return nil
}

// InputValidationDefinitions defines the validation of an input
type InputValidationDefinition struct {
	// DefaultValue is the default value of the variable
	// +kubebuilder:validation:Optional
	DefaultValue string `json:"defaultValue,omitempty"`

	// MinLength is the minimum number or length of the variable if a string
	// +kubebuilder:validation:Optional
	MinLength *int `json:"minLength,omitempty"`

	// MaxLength is the maximum number or length of the variable if a string
	// +kubebuilder:validation:Optional
	MaxLength *int `json:"maxLength,omitempty"`

	// Pattern is the pattern of the variable if a string
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty"`

	// Required is true if the variable is required
	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty"`
}

func (ivd *InputValidationDefinition) IsRequired() bool {
	return ivd != nil && ivd.Required != nil && *ivd.Required
}

func (ivd *InputValidationDefinition) HasDefault() bool {
	if ivd == nil {
		return false
	}
	return ivd.DefaultValue != ""
}

// InputContext defines an input context for a component resource
// Can be extended to reference other resources e.g. secrets
type InputContext struct {
	// Inputs are the set of input values for a component
	// +kubebuilder:validation:Optional
	Values InputValues `json:"values,omitempty"`
	// Vars are the set of user variables for the component
	// +kubebuilder:validation:Optional
	Vars []Var `json:"vars,omitempty"`
	// SecretsRef is the name of a workspace secret that contains the user-provided secret values
	// for this component.
	// +kubebuilder:validation:Optional
	SecretsRef string `json:"secretsRef,omitempty"`
	// InputDeps are the set of dependencies used to resolve the inputs of this components.
	// +kubebuilder:validation:Optional
	InputDeps []InputDependency `json:"inputDeps,omitempty"`
}

// InputValues defines a list of input values
// +kubebuilder:validation:Type=array
type InputValues []InputValue

func (ivs InputValues) Get(name string) *InputValue {
	_, iv := ivs.GetWithIndex(name)
	return iv
}

func (ivs InputValues) GetWithIndex(name string) (int, *InputValue) {
	for i, iv := range ivs {
		if iv.Name == name {
			return i, &iv
		}
	}

	return -1, nil
}

// InputValue defines a the value for an input
type InputValue struct {
	// Name is the name of the input
	// this must match the name of a valid input definition
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// Value is the value of the input
	// Can be a string, yaml or a Wayfinder template
	// +kubebuilder:validation:Optional
	Value string `json:"value"`
}

func (iv *InputValue) HasValue() bool {
	return iv != nil
}

// Var defines a "user" variable
// Used as user input (e.g. to deployment jobs)
// NOT as direct inputs to resources (see inputs)
type Var struct {
	// Name is the name of the variable
	// Can be referenced in input values in templates
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// Value is the value of the variable
	// +kubebuilder:validation:Required
	Value string `json:"value"`
}

type Vars []Var

func (vs Vars) Get(name string) (Var, bool) {
	for _, v := range vs {
		if v.Name == name {
			return v, true
		}
	}

	return Var{}, false
}

func (vs Vars) Has(name string) bool {
	_, ok := vs.Get(name)
	return ok
}

type InputDependency struct {
	// DepName is the 'logical' name of the dependency. This is the name used to reference this
	// dependency in templates.
	// +kubebuilder:validation:Required
	DepName string `json:"depName"`
	// Kind is the kind of the resource meeting this dependency
	// +kubebuilder:validation:Required
	Kind string `json:"kind"`
	// ResourceName is the name of the concrete resource that meets this dependency
	// +kubebuilder:validation:Required
	ResourceName string `json:"resourceName"`
	// Workspace is the workspace of the resource meeting this dependency. If unspecified, the
	// workspace of the containing resource will be used.
	// +kubebuilder:validation:Optional
	Workspace WorkspaceKey `json:"workspace,omitempty"`
}

func (i *InputDependency) IsSet() bool {
	return i != nil && i.ResourceName != ""
}
