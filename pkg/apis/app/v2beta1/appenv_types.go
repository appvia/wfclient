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
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	corev1alpha1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
)

const AppEnvKind = "AppEnv"

func AppEnvGroupVersionKind() schema.GroupVersionKind {
	return schema.FromAPIVersionAndKind(GroupVersion.String(), AppEnvKind)
}

// +kubebuilder:webhook:name=appenvs.app.appvia.io,mutating=false,path=/validate/app.appvia.io/appenvs,verbs=create;update;delete,groups="app.appvia.io",resources=appenvs,versions=v2beta1,failurePolicy=fail,sideEffects=None,admissionReviewVersions=v1
// +kubebuilder:webhook:name=appenvs.app.appvia.io,mutating=true,path=/mutate/app.appvia.io/appenvs,verbs=create;update,groups="app.appvia.io",resources=appenvs,versions=v2beta1,failurePolicy=fail,sideEffects=NoneOnDryRun,admissionReviewVersions=v1

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppEnv represents a deployable environment for an application - i.e. a namespace for the
// application's usage.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=appenvs,scope=Namespaced,categories={wayfinder}
type AppEnv struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppEnvSpec   `json:"spec,omitempty"`
	Status AppEnvStatus `json:"status,omitempty"`
}

func (AppEnv) APIPath() string {
	return "appenvs"
}

func (c *AppEnv) Clone() corev1alpha1.Object {
	return c.DeepCopy()
}

func (c *AppEnv) CloneInto(o corev1alpha1.Object) {
	c.DeepCopyInto(o.(*AppEnv))
}

func (AppEnv) ListType() corev1alpha1.ObjectList {
	return &AppEnvList{}
}

// GetCommonStatus implements CommonStatusAware
func (c *AppEnv) GetCommonStatus() *corev1alpha1.CommonStatus {
	return &c.Status.CommonStatus
}

func (c *AppEnv) GetRef() corev1alpha1.AppEnvRef {
	return corev1alpha1.AppEnvRef{
		Workspace: corev1alpha1.Workspace(c),
		App:       c.Spec.Application,
		EnvName:   c.Spec.Name,
	}
}

func (c *AppEnv) GetAppEnvObjectName() string {
	// Assume appenvs are named APP-ENV (this is validated by server/apps/handlers/internal/appenv/validate.go)
	return fmt.Sprintf("%s-%s", c.Spec.Application, c.Spec.Name)
}

// HostedByMultitenantCluster will return true if this appenv is hosted by a cluster from a
// different workspace
func (c *AppEnv) HostedByMultitenantCluster() bool {
	return c.Namespace != c.Spec.ClusterRef.Namespace
}

// AppEnvSpec defines an environment for an application
// +k8s:openapi-gen=true
type AppEnvSpec struct {
	// Cloud defines which cloud provider this application is being developed for.
	// +kubebuilder:validation:Enum=aws;gcp;azure
	// +kubebuilder:validation:Required
	Cloud string `json:"cloud"`
	// Application is the name of the application that this environment belongs to. It must comprise
	// only alphanumeric characters, and be 1-10 characters long.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Application string `json:"application"`
	// Name is the unique (within the application) human-readable name for this environment. It must
	// comprise only alphanumeric characters, and be 1-20 characters long.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// Key is a unique (within the appliction), short DNS-compatible name for this environment. If
	// unspecified on creation a suitable value will be derived from the name. If specified, it will
	// be validated for uniqueness on entry.
	// +kubebuilder:validation:MaxLength=8
	// +kubebuilder:validation:Optional
	Key string `json:"key,omitempty"`
	// ClusterRef defines the cluster on which this application environment should be hosted, if this
	// environment has an associated cluster. If both this and CloudAccessConfigRef are populated,
	// the cluster must use the same cloud access config as the one referenced here.
	// +kubebuilder:validation:Optional
	ClusterRef corev1alpha1.Ownership `json:"clusterRef,omitempty"`
	// CloudAccessConfigRef is a reference to the cloud access configuration to use when building
	// resources for this environment.
	// +kubebuilder:validation:Optional
	CloudAccessConfigRef *corev1alpha1.CloudAccessConfigRef `json:"cloudAccessConfigRef,omitempty"`
	// Namespace is the requested name for the environment's namespace on the specified cluster.
	// If unpopulated, Wayfinder will auto-populate this with a sensible name on entry.
	// +kubebuilder:validation:Optional
	Namespace string `json:"namespace"`
	// Stage is the infrastructure stage to which this environment belongs
	// +kubebuilder:validation:Required
	Stage string `json:"stage"`
	// Order gives a numeric ordering of this environment, used to sort environments in a logical
	// sequence. If two environments for an app have the same order, their display order is
	// undefined and may change.
	// +kubebuilder:validation:Optional
	Order *int `json:"order,omitempty"`
	// Vars is a set of variables specific to this app environment. These variables can be used in
	// container component environment variables (see AppComponent spec.container.containers[].env),
	// cloud resource component input variables (see AppComponent spec.cloudResource.envVars), and
	// deployment/ingress annotations (see AppComponent spec.container.deploymentAnnotations and
	// spec.container.expose.ingressAnnotations).
	// +kubebuilder:validation:Type=array
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	// +kubebuilder:validation:Optional
	Vars AppEnvVars `json:"vars,omitempty"`
}

type AppEnvVars []AppEnvVar

func (v AppEnvVars) ToMap() map[string]string {
	res := map[string]string{}
	for _, aev := range v {
		res[aev.Name] = aev.Value
	}
	return res
}

func (v AppEnvVars) Get(key string) string {
	for _, aev := range v {
		if aev.Name == key {
			return aev.Value
		}
	}
	return ""
}

type AppEnvVar struct {
	// Name is the name of this variable, used to reference it in app component definitions.
	// +kubebuilder:validation:Pattern=`^[a-zA-Z][a-zA-Z0-9_-]+$`
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// Value is the value of the variable
	// +kubebuilder:validation:Required
	Value string `json:"value"`
}

// AppEnvStatus defines the status of an application environment
// +k8s:openapi-gen=true
type AppEnvStatus struct {
	corev1alpha1.CommonStatus `json:",inline"`
	// HostEnv provides details about where this app env is hosted
	HostEnv AppEnvHostEnv `json:"hostEnv,omitempty"`
	// DNSZone is the DNS zone which should be used for this environment.
	// +kubebuilder:validation:Optional
	DNSZone string `json:"dnsZone,omitempty"`
	// DNSEntryTemplate is the template to use to generate DNS entries for this environment.
	// +kubebuilder:validation:Optional
	DNSEntryTemplate string `json:"dnsEntryTemplate,omitempty"`
	// PrivateDNSZone is the private DNS zone which should be used for this environment.
	// +kubebuilder:validation:Optional
	PrivateDNSZone string `json:"privateDnsZone,omitempty"`
	// PrivateDNSEntryTemplate is the template to use to generate private DNS entries for this environment.
	// +kubebuilder:validation:Optional
	PrivateDNSEntryTemplate string `json:"privateDnsEntryTemplate,omitempty"`
	// CertIssuers are the certificate issuers which can be used in this app env
	// +kubebuilder:validation:Optional
	CertIssuers []string `json:"certIssuers,omitempty"`
	// IngressClasses are the ingress classes which can be used in the app env
	// +kubebuilder:validation:Optional
	IngressClasses []IngressClass `json:"ingressClasses,omitempty"`
	// Deployment shows the deployed status of the app to this environment. The deployment
	// status will be updated approximately once per minute, to get up to date status, call the
	// deploystatus subresource API of the appenv.
	// +kubebuilder:validation:Optional
	Deployment AppEnvDeploymentStatus `json:"deployment,omitempty"`
	// DeploymentLastChecked identifies when the deployment status of this app was last checked.
	// +kubebuilder:validation:Optional
	DeploymentLastChecked metav1.Time `json:"deploymentLastChecked,omitempty"`
}

type AppEnvHostEnv struct {
	// CloudAccessConfigRef is a reference to the cloud access config that the host cluster for this
	// environment is using.
	CloudAccessConfigRef corev1alpha1.CloudAccessConfigRef `json:"cloudAccessConfigRef,omitempty"`
	// AccountIdentifier is the identifier of the cloud account/project/subscription in which the
	// cluster hosting this environment is located.
	AccountIdentifier string `json:"accountIdentifier,omitempty"`

	// NamespaceClaimRef is a reference to the namespace claim for this environment
	NamespaceClaimRef corev1alpha1.NamespaceClaimRef `json:"namespaceClaimRef,omitempty"`

	// ClusterRef is a reference to the cluster hosting this environment
	ClusterRef corev1alpha1.ClusterRef `json:"clusterRef,omitempty"`
}

// Empty returns true if the cloud access config, namespace and cluster refs are all unpopulated
func (a AppEnvHostEnv) Empty() bool {
	return a.CloudAccessConfigRef.Empty() && a.NamespaceClaimRef.Empty() && a.ClusterRef.Empty()
}

type IngressClass struct {
	// Class is the name of the ingress class
	// +kubebuilder:validation:Required
	Class string `json:"class"`
	// Namespace is the namespace the ingress controller is in
	// +kubebuilder:validation:Optional
	Namespace string `json:"namespace,omitempty"`
}

type AppEnvDeploymentStatus struct {
	// Deployed will be true if one or more components are deployed to this environment
	// +kubebuilder:validation:Optional
	Deployed bool `json:"deployed,omitempty"`
}

// AppEnvList provides a list of application environments
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AppEnvList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppEnv `json:"items"`
}

func (AppEnvList) ObjectType() corev1alpha1.Object {
	return &AppEnv{}
}

func (l *AppEnvList) Clone() corev1alpha1.ObjectList {
	return l.DeepCopy()
}

func (l *AppEnvList) CloneInto(o corev1alpha1.ObjectList) {
	l.DeepCopyInto(o.(*AppEnvList))
}

func (l *AppEnvList) GetItems() []corev1alpha1.Object {
	res := make([]corev1alpha1.Object, len(l.Items))
	for i, item := range l.Items {
		ref := item
		res[i] = &ref
	}
	return res
}

func (l *AppEnvList) SetItems(objects []corev1alpha1.Object) {
	l.Items = []AppEnv{}
	for _, o := range objects {
		l.Items = append(l.Items, *(o.(*AppEnv)))
	}
}
