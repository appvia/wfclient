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

package v2beta1

import corev1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"

const AppDefinitionVersionV1 = "v1"

type AppDefinitionData struct {
	// APIVersion is the version of this syntax for defining apps.
	APIVersion string `json:"apiVersion"`
	// App is the name of the app
	App string `json:"app"`
	// Description provides extended information about this app, allowing users to understand what it is and how it can be used.
	Description string `json:"description,omitempty"`
	// DeploymentSet is a name to identify a distinct set of components to deploy. This allows you
	// to specify separate sets of components to deploy in different Wayfinder.yaml files, without
	// them interfering with each other.
	DeploymentSet string `json:"deploymentSet,omitempty"`
	// Vars is a list of variables that can be used in the app definition. This allows you to add
	// descriptions and defaults to variables.
	Vars VarDefinitions `json:"vars,omitempty"`
	// Secrets is a list of secrets that can be used in the app definition. This allows you to add
	// descriptions and defaults to secrets.
	Secrets SecretDefinitions `json:"secrets,omitempty"`
	// Components is the set of components that make up the app
	Components ComponentDefinitions `json:"components,omitempty"`
}

// VarDefinitions is a list of variables that can be used in the app definition. This allows you to add
// descriptions and defaults to variables.
type VarDefinitions []VarDefinition

// SecretDefinitions is a list of secrets that can be used in the app definition. This allows you to add
// descriptions and defaults to secrets.
type SecretDefinitions []SecretDefinition

func (v VarDefinitions) Get(name string) (VarDefinition, bool) {
	for _, v := range v {
		if v.Name == name {
			return v, true
		}
	}
	return VarDefinition{}, false
}

func (v VarDefinitions) Has(name string) bool {
	_, ok := v.Get(name)
	return ok
}

func (s SecretDefinitions) Get(name string) (SecretDefinition, bool) {
	for _, s := range s {
		if s.Name == name {
			return s, true
		}
	}
	return SecretDefinition{}, false
}

func (s SecretDefinitions) Has(name string) bool {
	_, ok := s.Get(name)
	return ok
}

// ComponentDefinitions is a map of component names to component definitions
type ComponentDefinitions map[string]*ComponentDefinition

// List returns all of the component definitions in an array. Note that the order of the array
// is specifically NOT deterministic, so sort the result before usage if ordering is important.
func (c ComponentDefinitions) List() []*ComponentDefinition {
	componentList := []*ComponentDefinition{}
	for k := range c {
		comp := c[k]
		if comp.Name == "" {
			comp.Name = k
		}
		componentList = append(componentList, comp)
	}
	return componentList
}

type VarDefinition struct {
	Name string `json:"name"`
	// Description is a description of the variable
	Description string `json:"description,omitempty"`
	// Default is the default value of the variable if not provided.
	Default string `json:"default,omitempty"`
	// Required indicates that this variable must be provided. If default is set, this has no
	// effect. If default is not set and required is true, an error will occur if the variable is
	// not provided.
	Required bool `json:"required,omitempty"`
	// Sensitive indicates that this variable is sensitive and should not be handled in plain text.
	Sensitive bool `json:"sensitive,omitempty"`
}

type SecretDefinition struct {
	Name string `json:"name"`
	// Description is a description of the secret
	Description string `json:"description,omitempty"`
	// Required indicates that this secret must be provided. If required is true, an error will
	// occur if the secret is not provided.
	Required bool `json:"required,omitempty"`
}

func (a *AppDefinitionData) GetRequiredChartFiles() []string {
	res := []string{}
	for _, component := range a.Components {
		if component.Helm != nil && component.Helm.ChartPath != "" {
			res = append(res, component.Helm.ChartPath)
		}
	}
	return res
}

type ComponentDefinition struct {
	Component `json:",inline"`
	// Switch provides a conditional set of alternatives for this component.
	Switch []SwitchableComponent `json:"switch,omitempty"`
}

type SwitchableComponent struct {
	Case      string `json:"case"`
	Component `json:",inline"`
}

type Component struct {
	// Name is the component's name. This must be the same as the key used to identify the component
	// in the Components map if provided. If omitted, it will be auto-populated.
	Name string `json:"name,omitempty"`
	// Deps is a list of dependencies from this component to other components
	Deps []string `json:"deps,omitempty"`
	// Type is the type of the component - Package or CloudResource
	Type ComponentType `json:"type,omitempty"`
	// Plan is the plan to build the component off. For a Package, this is a reference to the
	// Package, and for a CloudResource, this is a reference to the CloudResourcePlan.
	Plan corev1.PlanRef `json:"plan,omitempty"`
	// Helm provides the details of a user-provided Helm chart to deploy. Only valid when type is Package.
	Helm *Helm `json:"helm,omitempty"`
	// DeploymentRoles causes the deployment of this package to operate with an identity using the specified role or roles.
	// These roles must be present in the target cluster, and must be labelled appropriately as per the cluster plan's App Deployment Cluster Role Labels setting.
	// If this is unpopulated, the Default App Deployment Cluster Role defined on the cluster plan will be used.
	// Only valid for Package components
	DeploymentRoles []string `json:"deploymentRoles,omitempty"`
	// Inputs is a list of inputs for the component
	Inputs corev1.InputValues `json:"inputs,omitempty"`
	// Outputs defines a set of outputs to generate when this resource is deployed. These can then
	// be consumed by other components and outputs of the app definition.
	Outputs []OutputDefinition `json:"outputs,omitempty"`
	// WorkloadIdentity specifies a workload identity for this component
	WorkloadIdentity *WorkloadIdentity `json:"workloadIdentity,omitempty"`
}

type Input struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Helm struct {
	// ChartURL is a direct URL to a helm chart. If specified, repoURL, chartName and chartVersion must be empty
	// +kubebuilder:validation:Optional
	ChartURL string `json:"chartURL,omitempty"`
	// ChartPath is a directory containing a helm chart. If specified, repoURL, chartURL and chartVersion must be empty
	// +kubebuilder:validation:Optional
	ChartPath string `json:"chartPath,omitempty"`
	// RepoURL is a URL to a helm repository. Must be combined with chart and chartVersion. Cannot be combined with chartURL or chartPath.
	// +kubebuilder:validation:Optional
	RepoURL string `json:"repoURL,omitempty"`
	// ChartName is the name of the chart to deploy.
	// +kubebuilder:validation:Required
	ChartName string `json:"chart"`
	// ChartVersion is the version of the chart identified by chartName and repoURL to use. Cannot be combined with chartURL or chartPath.
	// +kubebuilder:validation:Optional
	ChartVersion string `json:"chartVersion,omitempty"`
	// ValuesTemplate is a template for the values.yaml to use when installing the Helm chart
	// +kubebuilder:validation:Optional
	ValuesTemplate string `json:"valuesTemplate,omitempty"`
	// AdditionalValuesFiles is a list of additional value files to apply when installing the Helm
	// chart. These will be applied after the values template, so can override values from the
	// template. Do not include secret data in these files as they are handled in plain text.
	// +kubebuilder:validation:Optional
	AdditionalValuesFiles []AdditionalHelmValuesFile `json:"additionalValuesFiles,omitempty"`
}

type AdditionalHelmValuesFile struct {
	// Path is the relative path to the value file. Do not include secret data in this file as it is
	// handled in plain text. Use 'resolveSecret' in the ValuesTemplate to provide secrets.
	// +kubebuilder:validation:Required
	Path string `json:"path"`
	// Optional is true if the value file is optional and should not be required for the deployment
	// to succeed. If this is not specified, deployment will fail if no file is provided.
	// +kubebuilder:validation:Optional
	Optional bool `json:"optional,omitempty"`
}

type OutputDefinition struct {
	corev1.OutputDefinition `json:",inline"`
	// ValueTemplate is the value of the output. This can use template functions to produce the
	// value, using the same template context as the other templateable fields.
	// +kubebuilder:validation:Optional
	ValueTemplate string `json:"valueTemplate,omitempty"`
}

type WorkloadIdentity struct {
	// IdentityOnly will create an cloud managed identity for this component with no access permissions
	// Do not specify any access permissions if this is true
	IdentityOnly bool `json:"identityOnly,omitempty"`
	// ServiceAccountName is the name of the service account in Kubernetes that will have access to managed cloud identity
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
	// Access provides a list of access permissions that are required for this component to work
	Access []WorkloadIdentityAccess `json:"access,omitempty"`
}

type WorkloadIdentityAccess struct {
	// To is the name of a component that this component needs access to
	To string `json:"to"`
	// Permission is the name of the specific access permission to add to the workload identity
	// This is the name of the access permission in a CloudResourcePlan
	Permission string `json:"permission"`
}
