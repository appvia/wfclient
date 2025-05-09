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

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.

var map_AdditionalHelmValuesFile = map[string]string{
	"path":     "Path is the relative path to the value file. Do not include secret data in this file as it is handled in plain text. Use 'resolveSecret' in the ValuesTemplate to provide secrets.",
	"optional": "Optional is true if the value file is optional and should not be required for the deployment to succeed. If this is not specified, deployment will fail if no file is provided.",
}

func (AdditionalHelmValuesFile) SwaggerDoc() map[string]string {
	return map_AdditionalHelmValuesFile
}

var map_AppDefinitionData = map[string]string{
	"apiVersion":    "APIVersion is the version of this syntax for defining apps.",
	"app":           "App is the name of the app",
	"description":   "Description provides extended information about this app, allowing users to understand what it is and how it can be used.",
	"deploymentSet": "DeploymentSet is a name to identify a distinct set of components to deploy. This allows you to specify separate sets of components to deploy in different Wayfinder.yaml files, without them interfering with each other.",
	"vars":          "Vars is a list of variables that can be used in the app definition. This allows you to add descriptions and defaults to variables.",
	"secrets":       "Secrets is a list of secrets that can be used in the app definition. This allows you to add descriptions and defaults to secrets.",
	"components":    "Components is the set of components that make up the app",
}

func (AppDefinitionData) SwaggerDoc() map[string]string {
	return map_AppDefinitionData
}

var map_Component = map[string]string{
	"name":             "Name is the component's name. This must be the same as the key used to identify the component in the Components map if provided. If omitted, it will be auto-populated.",
	"deps":             "Deps is a list of dependencies from this component to other components",
	"type":             "Type is the type of the component - Package or CloudResource",
	"plan":             "Plan is the plan to build the component off. For a Package, this is a reference to the Package, and for a CloudResource, this is a reference to the CloudResourcePlan.",
	"helm":             "Helm provides the details of a user-provided Helm chart to deploy. Only valid when type is Package.",
	"deploymentRoles":  "DeploymentRoles causes the deployment of this package to operate with an identity using the specified role or roles. These roles must be present in the target cluster, and must be labelled appropriately as per the cluster plan's App Deployment Cluster Role Labels setting. If this is unpopulated, the Default App Deployment Cluster Role defined on the cluster plan will be used. Only valid for Package components",
	"inputs":           "Inputs is a list of inputs for the component",
	"outputs":          "Outputs defines a set of outputs to generate when this resource is deployed. These can then be consumed by other components and outputs of the app definition.",
	"workloadIdentity": "WorkloadIdentity specifies a workload identity for this component",
}

func (Component) SwaggerDoc() map[string]string {
	return map_Component
}

var map_ComponentDefinition = map[string]string{
	"switch": "Switch provides a conditional set of alternatives for this component.",
}

func (ComponentDefinition) SwaggerDoc() map[string]string {
	return map_ComponentDefinition
}

var map_Helm = map[string]string{
	"chartURL":              "ChartURL is a direct URL to a helm chart. If specified, repoURL, chartName and chartVersion must be empty",
	"chartPath":             "ChartPath is a directory containing a helm chart. If specified, repoURL, chartURL and chartVersion must be empty",
	"repoURL":               "RepoURL is a URL to a helm repository. Must be combined with chart and chartVersion. Cannot be combined with chartURL or chartPath.",
	"chart":                 "ChartName is the name of the chart to deploy.",
	"chartVersion":          "ChartVersion is the version of the chart identified by chartName and repoURL to use. Cannot be combined with chartURL or chartPath.",
	"valuesTemplate":        "ValuesTemplate is a template for the values.yaml to use when installing the Helm chart",
	"additionalValuesFiles": "AdditionalValuesFiles is a list of additional value files to apply when installing the Helm chart. These will be applied after the values template, so can override values from the template. Do not include secret data in these files as they are handled in plain text.",
}

func (Helm) SwaggerDoc() map[string]string {
	return map_Helm
}

var map_OutputDefinition = map[string]string{
	"valueTemplate": "ValueTemplate is the value of the output. This can use template functions to produce the value, using the same template context as the other templateable fields.",
}

func (OutputDefinition) SwaggerDoc() map[string]string {
	return map_OutputDefinition
}

var map_SecretDefinition = map[string]string{
	"description": "Description is a description of the secret",
	"required":    "Required indicates that this secret must be provided. If required is true, an error will occur if the secret is not provided.",
}

func (SecretDefinition) SwaggerDoc() map[string]string {
	return map_SecretDefinition
}

var map_VarDefinition = map[string]string{
	"description": "Description is a description of the variable",
	"default":     "Default is the default value of the variable if not provided.",
	"required":    "Required indicates that this variable must be provided. If default is set, this has no effect. If default is not set and required is true, an error will occur if the variable is not provided.",
	"sensitive":   "Sensitive indicates that this variable is sensitive and should not be handled in plain text.",
}

func (VarDefinition) SwaggerDoc() map[string]string {
	return map_VarDefinition
}

var map_WorkloadIdentity = map[string]string{
	"identityOnly":       "IdentityOnly will create an cloud managed identity for this component with no access permissions Do not specify any access permissions if this is true",
	"serviceAccountName": "ServiceAccountName is the name of the service account in Kubernetes that will have access to managed cloud identity",
	"access":             "Access provides a list of access permissions that are required for this component to work",
}

func (WorkloadIdentity) SwaggerDoc() map[string]string {
	return map_WorkloadIdentity
}

var map_WorkloadIdentityAccess = map[string]string{
	"to":         "To is the name of a component that this component needs access to",
	"permission": "Permission is the name of the specific access permission to add to the workload identity This is the name of the access permission in a CloudResourcePlan",
}

func (WorkloadIdentityAccess) SwaggerDoc() map[string]string {
	return map_WorkloadIdentityAccess
}

var map_AppDeploymentJob = map[string]string{
	"": "AppDeploymentJob represents a deployment job / action on one or more application components.",
}

func (AppDeploymentJob) SwaggerDoc() map[string]string {
	return map_AppDeploymentJob
}

var map_AppDeploymentJobComponentStatus = map[string]string{
	"name":          "Name is the name of the component",
	"status":        "Status of the component",
	"message":       "Message is the status message for the component",
	"errorDetail":   "ErrorDetail is the summary for why the component failed",
	"conditions":    "Conditions from the component",
	"resourceKind":  "ResourceKind identifies the deployed type of the Wayfinder resource",
	"resourceName":  "ResourceName identifies the deployed name of the Wayfinder resource in the workspace of this deployment job",
	"dnsEntries":    "DNSEntries is the set of DNS entries that have been generated for this component by usage of the getDNSEntry, getPrivateDNSEntry and getPublicDNSEntry template functions",
	"outputs":       "Outputs provide the outputs generated by this component. The values of outputs marked Sensitive will not be populated - see OutputsSecret for a reference to the workspace secret which contains all output values including those marked sensitive. not included.",
	"outputsSecret": "OutputsSecret is the name of a secret in the workspace of this job that contains all output values for this component, including those marked sensitive. Note that the secret may have been further updated since this job ran.",
}

func (AppDeploymentJobComponentStatus) SwaggerDoc() map[string]string {
	return map_AppDeploymentJobComponentStatus
}

var map_AppDeploymentJobList = map[string]string{
	"": "AppDeploymentJobList provides a list of AppDeploymentJobs",
}

func (AppDeploymentJobList) SwaggerDoc() map[string]string {
	return map_AppDeploymentJobList
}

var map_AppDeploymentJobSpec = map[string]string{
	"":              "AppDeploymentJobSpec defines the desired state of AppDeploymentJob",
	"appDefinition": "AppDefinition is the definition of the application to deploy",
	"appEnvRef":     "AppEnvRef is the reference to the application environment to deploy",
	"deploymentSet": "DeploymentSet is the name of the deployment set to use for this deployment job. This causes the removal of other components in the same deploymentset if they are not included in the components array. If unset, a 'default' deployment set is used.",
	"only":          "Only limits the set of components that should be deployed. If unset, all components in the deployment set are deployed.",
	"action":        "Action is the type of deployment action",
	"vars":          "Vars are the set of user variables for the deployment job Represent the values that can be used in the component input value templates",
	"secretsRef":    "SecretsRef is the name of a workspace secret that contains the user-provided secret values for this deployment job.",
	"files":         "Files is a list of file data to be used to perform the deployment. These are referenced from component definitions by path and type.",
	"refresh":       "Refresh ensures re-deployment of the component(s) even if they already exist and are unchanged if this value is different to any previous refresh value on the components. This is typically set to the timestamp that the refresh was requested at.",
	"customTimeout": "CustomTimeout is the optional timeout for the deployment job. If this is not provided, a system default timeout will be used. After this period, the deployment will fail if it has not successfully completed.",
}

func (AppDeploymentJobSpec) SwaggerDoc() map[string]string {
	return map_AppDeploymentJobSpec
}

var map_AppDeploymentJobStatus = map[string]string{
	"":                    "AppDeploymentJobStatus defines the observed state of AppDeploymentJob",
	"plan":                "Plan is the set of actions that are being taken by this deployment job",
	"components":          "Components are the status of the components",
	"inProgressComponent": "InProgressComponent is the name of the component that is currently being deployed by this job. It will be empty if no component is currently being deployed.",
	"completedAt":         "CompletedAt will be populated once this job has completed processing.",
}

func (AppDeploymentJobStatus) SwaggerDoc() map[string]string {
	return map_AppDeploymentJobStatus
}

var map_AppDeploymentPlan = map[string]string{
	"actions": "Actions is the set of actions that are being taken by this deployment job",
}

func (AppDeploymentPlan) SwaggerDoc() map[string]string {
	return map_AppDeploymentPlan
}

var map_AppDeploymentPlanAction = map[string]string{
	"component":     "Component is the name of the component that is being deployed",
	"componentType": "ComponentType is the type of the component",
	"resourceKind":  "ResourceKind is the kind of the resource that is being deployed/removed.",
	"resourceName":  "ResourceName is the name of the resource that is being deployed/removed.",
	"action":        "Action is the type of deployment action",
	"status":        "Status is the status of the action",
}

func (AppDeploymentPlanAction) SwaggerDoc() map[string]string {
	return map_AppDeploymentPlanAction
}

var map_AppDeploymentFile = map[string]string{
	"type": "Type is the type of file to be added.",
	"path": "Path is the path to the file as used in the component definition.",
	"data": "Data is the data of the file. For Helm chart files, this should be a base64 encoded tar.gz. For Helm values files, this should be the raw YAML data.",
}

func (AppDeploymentFile) SwaggerDoc() map[string]string {
	return map_AppDeploymentFile
}

var map_AppDeploymentRequest = map[string]string{
	"appManifest":       "AppManifest is the provided JSON/YAML manifest from the user",
	"appManifestFormat": "AppManifestFormat is the format of the provided manifest ('json' or 'yaml'; default is 'yaml')",
	"appDefinitionRef":  "AppDefinitionRef is the reference to the AppDefinition to use for the deployment instead of the AppManifest.",
	"Vars":              "Vars is the set of deploy-time variables provided with the deployment request",
	"Secrets":           "Secrets is the set of deploy-time sensitive variables provided with the deployment request",
	"targetEnv":         "TargetEnv is the name of the target appenv",
	"customTimeout":     "CustomTimeout is the optional timeout for the deployment job. If this is not provided, a system default timeout will be used. After this period, the deployment will fail if it has not successfully completed.",
	"only":              "Only limits the deployment/removal to the specified components",
	"asMe":              "AsMe indicates if we should run as the current user instead of using deployment roles",
	"files":             "Files is a list of files to be used to perform the deployment. For each component that has a helm.chartPath populated, a HelmChart file should be added to this array. For each entry in helm.additionalValuesFiles populated in each component, a HelmValuesFile should be added to this array.",
}

func (AppDeploymentRequest) SwaggerDoc() map[string]string {
	return map_AppDeploymentRequest
}

var map_AppDefinition = map[string]string{
	"": "AppDefinition provides a definition of an app recorded by Wayfinder, allowing deployment of an app without a local Wayfinder.yaml file",
}

func (AppDefinition) SwaggerDoc() map[string]string {
	return map_AppDefinition
}

var map_AppDefinitionList = map[string]string{
	"": "AppDefinitionList provides a list of app definitions",
}

func (AppDefinitionList) SwaggerDoc() map[string]string {
	return map_AppDefinitionList
}

var map_AppDefinitionSpec = map[string]string{
	"":           "AppDefinitionSpec defines the specification for an AppDefinition",
	"version":    "Version is the version of the app definition",
	"definition": "Definition contains the actual AppDefinition details",
	"files":      "Files is a list of files to be used to perform the deployment. For each component that has a helm.chartPath populated, a HelmChart file should be added to this array. For each entry in helm.additionalValuesFiles populated in each component, a HelmValuesFile should be added to this array.",
}

func (AppDefinitionSpec) SwaggerDoc() map[string]string {
	return map_AppDefinitionSpec
}

var map_AppDefinitionStatus = map[string]string{
	"": "AppDefinitionStatus defines the status for an AppDefinition",
}

func (AppDefinitionStatus) SwaggerDoc() map[string]string {
	return map_AppDefinitionStatus
}

var map_AppEnv = map[string]string{
	"": "AppEnv represents a deployable environment for an application - i.e. a namespace for the application's usage.",
}

func (AppEnv) SwaggerDoc() map[string]string {
	return map_AppEnv
}

var map_AppEnvDeploymentStatus = map[string]string{
	"deployed": "Deployed will be true if one or more components are deployed to this environment",
}

func (AppEnvDeploymentStatus) SwaggerDoc() map[string]string {
	return map_AppEnvDeploymentStatus
}

var map_AppEnvHostEnv = map[string]string{
	"cloudAccessConfigRef": "CloudAccessConfigRef is a reference to the cloud access config that the host cluster for this environment is using.",
	"accountIdentifier":    "AccountIdentifier is the identifier of the cloud account/project/subscription in which the cluster hosting this environment is located.",
	"namespaceClaimRef":    "NamespaceClaimRef is a reference to the namespace claim for this environment",
	"clusterRef":           "ClusterRef is a reference to the cluster hosting this environment",
}

func (AppEnvHostEnv) SwaggerDoc() map[string]string {
	return map_AppEnvHostEnv
}

var map_AppEnvList = map[string]string{
	"": "AppEnvList provides a list of application environments",
}

func (AppEnvList) SwaggerDoc() map[string]string {
	return map_AppEnvList
}

var map_AppEnvSpec = map[string]string{
	"":                     "AppEnvSpec defines an environment for an application",
	"cloud":                "Cloud defines which cloud provider this application is being developed for.",
	"application":          "Application is the name of the application that this environment belongs to. It must comprise only alphanumeric characters, and be 1-10 characters long.",
	"name":                 "Name is the unique (within the application) human-readable name for this environment. It must comprise only alphanumeric characters, and be 1-20 characters long.",
	"key":                  "Key is a unique (within the appliction), short DNS-compatible name for this environment. If unspecified on creation a suitable value will be derived from the name. If specified, it will be validated for uniqueness on entry.",
	"clusterRef":           "ClusterRef defines the cluster on which this application environment should be hosted, if this environment has an associated cluster. If both this and CloudAccessConfigRef are populated, the cluster must use the same cloud access config as the one referenced here.",
	"cloudAccessConfigRef": "CloudAccessConfigRef is a reference to the cloud access configuration to use when building resources for this environment.",
	"namespace":            "Namespace is the requested name for the environment's namespace on the specified cluster. If unpopulated, Wayfinder will auto-populate this with a sensible name on entry.",
	"stage":                "Stage is the infrastructure stage to which this environment belongs",
	"order":                "Order gives a numeric ordering of this environment, used to sort environments in a logical sequence. If two environments for an app have the same order, their display order is undefined and may change.",
	"vars":                 "Vars is a set of variables specific to this app environment. These variables can be used in container component environment variables (see AppComponent spec.container.containers[].env), cloud resource component input variables (see AppComponent spec.cloudResource.envVars), and deployment/ingress annotations (see AppComponent spec.container.deploymentAnnotations and spec.container.expose.ingressAnnotations).",
}

func (AppEnvSpec) SwaggerDoc() map[string]string {
	return map_AppEnvSpec
}

var map_AppEnvStatus = map[string]string{
	"":                        "AppEnvStatus defines the status of an application environment",
	"hostEnv":                 "HostEnv provides details about where this app env is hosted",
	"dnsZone":                 "DNSZone is the DNS zone which should be used for this environment.",
	"dnsEntryTemplate":        "DNSEntryTemplate is the template to use to generate DNS entries for this environment.",
	"privateDnsZone":          "PrivateDNSZone is the private DNS zone which should be used for this environment.",
	"privateDnsEntryTemplate": "PrivateDNSEntryTemplate is the template to use to generate private DNS entries for this environment.",
	"certIssuers":             "CertIssuers are the certificate issuers which can be used in this app env",
	"ingressClasses":          "IngressClasses are the ingress classes which can be used in the app env",
	"deployment":              "Deployment shows the deployed status of the app to this environment. The deployment status will be updated approximately once per minute, to get up to date status, call the deploystatus subresource API of the appenv.",
	"deploymentLastChecked":   "DeploymentLastChecked identifies when the deployment status of this app was last checked.",
}

func (AppEnvStatus) SwaggerDoc() map[string]string {
	return map_AppEnvStatus
}

var map_AppEnvVar = map[string]string{
	"name":  "Name is the name of this variable, used to reference it in app component definitions.",
	"value": "Value is the value of the variable",
}

func (AppEnvVar) SwaggerDoc() map[string]string {
	return map_AppEnvVar
}

var map_IngressClass = map[string]string{
	"class":     "Class is the name of the ingress class",
	"namespace": "Namespace is the namespace the ingress controller is in",
}

func (IngressClass) SwaggerDoc() map[string]string {
	return map_IngressClass
}
