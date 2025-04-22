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

	corev1alpha1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
)

const AppDefinitionKind = "AppDefinition"

func AppDefinitionGroupVersionKind() schema.GroupVersionKind {
	return schema.FromAPIVersionAndKind(GroupVersion.String(), AppDefinitionKind)
}

// +kubebuilder:webhook:name=appdefinitions.app.appvia.io,mutating=false,path=/validate/app.appvia.io/appdefinitions,verbs=create;update;delete,groups="app.appvia.io",resources=appdefinitions,versions=v2beta1,failurePolicy=fail,sideEffects=None,admissionReviewVersions=v1
// +kubebuilder:webhook:name=appdefinitions.app.appvia.io,mutating=true,path=/mutate/app.appvia.io/appdefinitions,verbs=create;update,groups="app.appvia.io",resources=appdefinitions,versions=v2beta1,failurePolicy=fail,sideEffects=NoneOnDryRun,admissionReviewVersions=v1

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppDefinition provides a definition of an app recorded by Wayfinder, allowing deployment of an app without a local Wayfinder.yaml file
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=appdefinitions,scope=Cluster,categories={wayfinder}
type AppDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppDefinitionSpec   `json:"spec,omitempty"`
	Status AppDefinitionStatus `json:"status,omitempty"`
}

func (AppDefinition) APIPath() string {
	return "appdefinitions"
}

func (c *AppDefinition) Clone() corev1alpha1.Object {
	return c.DeepCopy()
}

func (c *AppDefinition) CloneInto(o corev1alpha1.Object) {
	c.DeepCopyInto(o.(*AppDefinition))
}

func (AppDefinition) ListType() corev1alpha1.ObjectList {
	return &AppDefinitionList{}
}

// GetCommonStatus implements CommonStatusAware
func (c *AppDefinition) GetCommonStatus() *corev1alpha1.CommonStatus {
	return &c.Status.CommonStatus
}

// AppDefinitionSpec defines the specification for an AppDefinition
// +k8s:openapi-gen=true
type AppDefinitionSpec struct {
	// Version is the version of the app definition
	Version corev1alpha1.ObjectVersion `json:"version,omitempty"`
	// Definition contains the actual AppDefinition details
	Definition AppDefinitionData `json:"definition"`
	// Files is a list of files to be used to perform the deployment. For each component that has
	// a helm.chartPath populated, a HelmChart file should be added to this array. For each entry
	// in helm.additionalValuesFiles populated in each component, a HelmValuesFile should be added to
	// this array.
	Files []AppDeploymentFile `json:"files,omitempty"`
}

// AppDefinitionStatus defines the status for an AppDefinition
// +k8s:openapi-gen=true
type AppDefinitionStatus struct {
	corev1alpha1.CommonStatus `json:",inline"`
}

// AppDefinitionList provides a list of app definitions
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AppDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppDefinition `json:"items"`
}

func (AppDefinitionList) ObjectType() corev1alpha1.Object {
	return &AppDefinition{}
}

func (l *AppDefinitionList) Clone() corev1alpha1.ObjectList {
	return l.DeepCopy()
}

func (l *AppDefinitionList) CloneInto(o corev1alpha1.ObjectList) {
	l.DeepCopyInto(o.(*AppDefinitionList))
}

func (l *AppDefinitionList) GetItems() []corev1alpha1.Object {
	res := make([]corev1alpha1.Object, len(l.Items))
	for i, item := range l.Items {
		ref := item
		res[i] = &ref
	}
	return res
}

func (l *AppDefinitionList) SetItems(objects []corev1alpha1.Object) {
	l.Items = []AppDefinition{}
	for _, o := range objects {
		l.Items = append(l.Items, *(o.(*AppDefinition)))
	}
}

// Implement corev1.Versioned interface
func (c *AppDefinition) VersionOf() string {
	return corev1alpha1.GetVersionedObjectName(c)
}

func (c *AppDefinition) GetVersion() corev1alpha1.ObjectVersion {
	return c.Spec.Version
}

func (c *AppDefinition) SetVersion(v corev1alpha1.ObjectVersion) {
	c.Spec.Version = v
}

func (c *AppDefinition) SetTags(tags []string) {
	corev1alpha1.SetObjectTags(c, tags)
}

func (c *AppDefinition) SetDescription(description string) {
	c.Spec.Definition.Description = description
}
