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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	corev1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
)

const AppDeploymentJobKind = "AppDeploymentJob"

func AppDeploymentJobGroupVersionKind() schema.GroupVersionKind {
	return schema.FromAPIVersionAndKind(GroupVersion.String(), AppDeploymentJobKind)
}

// +kubebuilder:webhook:name=appdeploymentjobs.app.appvia.io,mutating=false,path=/validate/app.appvia.io/appdeploymentjobs,verbs=create;update;delete,groups="app.appvia.io",resources=appdeploymentjobs,versions=v2beta1,failurePolicy=fail,sideEffects=None,admissionReviewVersions=v1
// +kubebuilder:webhook:name=appdeploymentjobs.app.appvia.io,mutating=true,path=/mutate/app.appvia.io/appdeploymentjobs,verbs=create;update,groups="app.appvia.io",resources=appdeploymentjobs,versions=v2beta1,failurePolicy=fail,sideEffects=NoneOnDryRun,admissionReviewVersions=v1

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppDeploymentJob represents a deployment job / action on one or more application components.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=appdeploymentjobs,scope=Namespaced,categories={wayfinder}
type AppDeploymentJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppDeploymentJobSpec   `json:"spec,omitempty"`
	Status AppDeploymentJobStatus `json:"status,omitempty"`
}

func (AppDeploymentJob) APIPath() string {
	return "appdeploymentjobs"
}

func (a *AppDeploymentJob) Clone() corev1.Object {
	return a.DeepCopy()
}

func (a *AppDeploymentJob) CloneInto(o corev1.Object) {
	a.DeepCopyInto(o.(*AppDeploymentJob))
}

func (AppDeploymentJob) ListType() corev1.ObjectList {
	return &AppDeploymentJobList{}
}

// GetCommonStatus implements CommonStatusAware for this resource
func (a *AppDeploymentJob) GetCommonStatus() *corev1.CommonStatus {
	return &a.Status.CommonStatus
}

// DeploymentAction defines the type of deployment action
// +kubebuilder:validation:Enum=Deploy;Remove
type DeploymentAction string

const (
	DeploymentActionDeploy DeploymentAction = "Deploy"
	DeploymentActionRemove DeploymentAction = "Remove"
)

// AppDeploymentJobSpec defines the desired state of AppDeploymentJob
// +k8s:openapi-gen=true
type AppDeploymentJobSpec struct {
	// AppDefinition is the definition of the application to deploy
	// +kubebuilder:validation:Required
	AppDefinition AppDefinitionData `json:"appDefinition"`

	// AppEnvRef is the reference to the application environment to deploy
	// +kubebuilder:validation:Required
	AppEnvRef corev1.AppEnvRef `json:"appEnvRef"`

	// CreateAppEnv provides options to create the referenced application environment if it does
	// not already exist. The name of the desired application environment must be specified in the
	// AppEnvRef.EnvName. If the environment already exists, populating this field will cause a
	// validation error.
	// +kubebuilder:validation:Optional
	CreateAppEnv *CreateAppEnvOptions `json:"createAppEnv,omitempty"`

	// DeploymentSet is the name of the deployment set to use for this deployment job. This causes
	// the removal of other components in the same deploymentset if they are not included in the
	// components array. If unset, a 'default' deployment set is used.
	// +kubebuilder:validation:Optional
	DeploymentSet string `json:"deploymentSet,omitempty"`
	// Only limits the set of components that should be deployed. If unset, all components in the
	// deployment set are deployed.
	// +kubebuilder:validation:Optional
	Only []string `json:"only,omitempty"`
	// Action is the type of deployment action
	// +kubebuilder:validation:Required
	Action DeploymentAction `json:"action"`
	// Vars are the set of user variables for the deployment job
	// Represent the values that can be used in the component input value templates
	// +kubebuilder:validation:Optional
	Vars []corev1.Var `json:"vars,omitempty"`
	// SecretsRef is the name of a workspace secret that contains the user-provided secret values
	// for this deployment job.
	// +kubebuilder:validation:Optional
	SecretsRef string `json:"secretsRef,omitempty"`

	// Files is a list of file data to be used to perform the deployment. These are referenced from
	// component definitions by path and type.
	// +kubebuilder:validation:Optional
	Files AppDeploymentFiles `json:"files,omitempty"`

	// Refresh ensures re-deployment of the component(s) even if they already exist and are
	// unchanged if this value is different to any previous refresh value on the components.
	// This is typically set to the timestamp that the refresh was requested at.
	// +kubebuilder:validation:Optional
	Refresh string `json:"refresh,omitempty"`

	// CustomTimeout is the optional timeout for the deployment job. If this is not provided, a
	// system default timeout will be used. After this period, the deployment will fail if it has
	// not successfully completed.
	// +kubebuilder:validation:Optional
	CustomTimeout metav1.Duration `json:"customTimeout,omitempty"`
}

type CreateAppEnvOptions struct {
	// Create should be set to true to build the app env if it does not exist. If this is not true,
	// the job will fail if the app env does not exist.
	// +kubebuilder:validation:Required
	Create bool `json:"create"`
	// CloudAccessConfigRef is the reference to the cloud access config to use to create the application
	// environment. Must specify either this or ClusterRef.
	// +kubebuilder:validation:Optional
	CloudAccessConfigRef *corev1.CloudAccessConfigRef `json:"cloudAccessConfigRef,omitempty"`
	// ClusterRef is the reference to the cluster to use to create the application environment.
	// Must specify either this or CloudAccessConfigRef.
	// +kubebuilder:validation:Optional
	ClusterRef *corev1.ClusterRef `json:"clusterRef,omitempty"`
}

// HasPackages returns true if the job has any package components
func (s *AppDeploymentJobSpec) HasPackages() bool {
	for _, component := range s.AppDefinition.Components {
		if component.Type == ComponentTypeKubeResource {
			return true
		}
	}
	return false
}

func (s *AppDeploymentJobSpec) GetComponent(name string) *ComponentDefinition {
	return s.AppDefinition.Components[name]
}

func (s *AppDeploymentJobSpec) GetComponentNames() []string {
	res := []string{}
	for name := range s.AppDefinition.Components {
		res = append(res, name)
	}
	return res
}

// Components returns all of the component definitions in an array. Note that the order of the array
// is specifically NOT deterministic, so sort the result before usage if ordering is important.
func (s *AppDeploymentJobSpec) Components() []*ComponentDefinition {
	componentList := []*ComponentDefinition{}
	for k := range s.AppDefinition.Components {
		comp := s.AppDefinition.Components[k]
		if comp.Name == "" {
			comp.Name = k
		}
		componentList = append(componentList, comp)
	}
	return componentList
}

const DefaultDeploymentSet = "default"

func (s *AppDeploymentJobSpec) GetDeploymentSet() string {
	if s.DeploymentSet != "" {
		return s.DeploymentSet
	}
	return DefaultDeploymentSet
}

// ComponentType defines the type of an application component
// +kubebuilder:validation:Enum=CloudResource;KubeResource;Package
type ComponentType string

const (
	ComponentTypeCloudResource ComponentType = "CloudResource"
	ComponentTypeKubeResource  ComponentType = "KubeResource"
)

func (w *WorkloadIdentity) Enabled() bool {
	return w != nil && (w.IdentityOnly || len(w.Access) > 0)
}

func (c *Component) RequiresPlan() bool {
	if c == nil {
		return false
	}
	// we always require a plan unless it's a helm chart
	return c.Helm == nil && c.Type == ComponentTypeKubeResource
}

// AppDeploymentJobStatus defines the observed state of AppDeploymentJob
// +k8s:openapi-gen=true
type AppDeploymentJobStatus struct {
	corev1.CommonStatus `json:",inline"`
	// Plan is the set of actions that are being taken by this deployment job
	// +kubebuilder:validation:Optional
	Plan *AppDeploymentPlan `json:"plan,omitempty"`
	// Components are the status of the components
	// +kubebuilder:validation:Optional
	Components []AppDeploymentJobComponentStatus `json:"components,omitempty"`
	// InProgressComponent is the name of the component that is currently being deployed by this
	// job. It will be empty if no component is currently being deployed.
	// +kubebuilder:validation:Optional
	InProgressComponent string `json:"inProgressComponent,omitempty"`
	// CompletedAt will be populated once this job has completed processing.
	// +kubebuilder:validation:Optional
	CompletedAt *metav1.Time `json:"completedAt,omitempty"`
}

type AppDeploymentPlan struct {
	// Actions is the set of actions that are being taken by this deployment job
	// +kubebuilder:validation:Optional
	Actions []AppDeploymentPlanAction `json:"actions,omitempty"`
}

type AppDeploymentPlanAction struct {
	// Component is the name of the component that is being deployed
	// +kubebuilder:validation:Required
	Component string `json:"component"`
	// ComponentType is the type of the component
	// +kubebuilder:validation:Required
	ComponentType ComponentType `json:"componentType"`
	// ResourceKind is the kind of the resource that is being deployed/removed.
	// +kubebuilder:validation:Optional
	ResourceKind metav1.GroupVersionKind `json:"resourceKind,omitempty"`
	// ResourceName is the name of the resource that is being deployed/removed.
	// +kubebuilder:validation:Optional
	ResourceName string `json:"resourceName,omitempty"`
	// Action is the type of deployment action
	// +kubebuilder:validation:Required
	Action DeploymentPlanAction `json:"action"`
	// Status is the status of the action
	// +kubebuilder:validation:Required
	Status DeploymentPlanActionStatus `json:"status"`
}

type DeploymentPlanActionStatus string

const (
	DeploymentPlanActionStatusPending    DeploymentPlanActionStatus = "Pending"
	DeploymentPlanActionStatusApplying   DeploymentPlanActionStatus = "Applying"
	DeploymentPlanActionStatusInProgress DeploymentPlanActionStatus = "InProgress"
	DeploymentPlanActionStatusComplete   DeploymentPlanActionStatus = "Complete"
	DeploymentPlanActionStatusFailed     DeploymentPlanActionStatus = "Failed"
)

// DeploymentPlanAction defines the type of action that a single step in the deployment plan will
// take.
// +kubebuilder:validation:Enum=Create;Update;Delete
type DeploymentPlanAction string

const (
	DeploymentPlanActionCreate DeploymentPlanAction = "Create"
	DeploymentPlanActionUpdate DeploymentPlanAction = "Update"
	DeploymentPlanActionDelete DeploymentPlanAction = "Delete"
)

// GetInProgressComponent returns the currently-in-progress component or nil if nothing is currently
// in progress on this job.
func (s *AppDeploymentJobStatus) GetInProgressComponent() *AppDeploymentJobComponentStatus {
	for _, c := range s.Components {
		if c.Name == s.InProgressComponent {
			return &c
		}
	}
	return nil
}

type AppDeploymentJobComponentStatus struct {
	// Name is the name of the component
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// Status of the component
	// +kubebuilder:validation:Required
	Status corev1.Status `json:"status,omitempty"`
	// Message is the status message for the component
	Message string `json:"message,omitempty"`
	// ErrorDetail is the summary for why the component failed
	ErrorDetail string `json:"errorDetail,omitempty"`
	// Conditions from the component
	Conditions []corev1.Condition `json:"conditions,omitempty"`
	// ResourceKind identifies the deployed type of the Wayfinder resource
	ResourceKind string `json:"resourceKind,omitempty"`
	// ResourceName identifies the deployed name of the Wayfinder resource in the workspace of this
	// deployment job
	ResourceName string `json:"resourceName,omitempty"`
	// DNSEntries is the set of DNS entries that have been generated for this component by
	// usage of the getDNSEntry, getPrivateDNSEntry and getPublicDNSEntry template functions
	// +kubebuilder:validation:Optional
	DNSEntries []corev1.DNSEntry `json:"dnsEntries,omitempty"`
	// Outputs provide the outputs generated by this component. The values of outputs marked
	// Sensitive will not be populated - see OutputsSecret for a reference to the workspace secret
	// which contains all output values including those marked sensitive.
	// not included.
	Outputs []corev1.Output `json:"outputs,omitempty"`
	// OutputsSecret is the name of a secret in the workspace of this job that contains all output
	// values for this component, including those marked sensitive. Note that the secret may have
	// been further updated since this job ran.
	OutputsSecret string `json:"outputsSecret,omitempty"`
}

// GetDNSEntries implements DNSEntryAware for AppDeploymentJob
func (a *AppDeploymentJob) GetDNSEntries() []corev1.DNSEntry {
	entries := []corev1.DNSEntry{}
	for i, c := range a.Status.Components {
		for j, d := range c.DNSEntries {
			exists := false
			for _, existingEntry := range entries {
				if existingEntry.Equals(&d) {
					exists = true
				}
			}
			if !exists {
				entries = append(entries, a.Status.Components[i].DNSEntries[j])
			}
		}
	}
	return entries
}

func (a *AppDeploymentJob) GetOutputs() []corev1.Output {
	outputs := []corev1.Output{}
	for _, c := range a.Status.Components {
		outputs = append(outputs, c.Outputs...)
	}
	return outputs
}

// AppDeploymentJobList provides a list of AppDeploymentJobs
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AppDeploymentJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppDeploymentJob `json:"items"`
}

func (AppDeploymentJobList) ObjectType() corev1.Object {
	return &AppDeploymentJob{}
}

func (l *AppDeploymentJobList) Clone() corev1.ObjectList {
	return l.DeepCopy()
}

func (l *AppDeploymentJobList) CloneInto(o corev1.ObjectList) {
	l.DeepCopyInto(o.(*AppDeploymentJobList))
}

func (l *AppDeploymentJobList) GetItems() []corev1.Object {
	res := make([]corev1.Object, len(l.Items))
	for i, item := range l.Items {
		ref := item
		res[i] = &ref
	}
	return res
}

func (l *AppDeploymentJobList) SetItems(objects []corev1.Object) {
	l.Items = []AppDeploymentJob{}
	for _, o := range objects {
		l.Items = append(l.Items, *(o.(*AppDeploymentJob)))
	}
}
