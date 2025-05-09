//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v2beta1

import (
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalHelmValuesFile) DeepCopyInto(out *AdditionalHelmValuesFile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalHelmValuesFile.
func (in *AdditionalHelmValuesFile) DeepCopy() *AdditionalHelmValuesFile {
	if in == nil {
		return nil
	}
	out := new(AdditionalHelmValuesFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDefinition) DeepCopyInto(out *AppDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDefinition.
func (in *AppDefinition) DeepCopy() *AppDefinition {
	if in == nil {
		return nil
	}
	out := new(AppDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDefinitionData) DeepCopyInto(out *AppDefinitionData) {
	*out = *in
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(VarDefinitions, len(*in))
		copy(*out, *in)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(SecretDefinitions, len(*in))
		copy(*out, *in)
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make(ComponentDefinitions, len(*in))
		for key, val := range *in {
			var outVal *ComponentDefinition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ComponentDefinition)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDefinitionData.
func (in *AppDefinitionData) DeepCopy() *AppDefinitionData {
	if in == nil {
		return nil
	}
	out := new(AppDefinitionData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDefinitionList) DeepCopyInto(out *AppDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDefinitionList.
func (in *AppDefinitionList) DeepCopy() *AppDefinitionList {
	if in == nil {
		return nil
	}
	out := new(AppDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDefinitionSpec) DeepCopyInto(out *AppDefinitionSpec) {
	*out = *in
	in.Definition.DeepCopyInto(&out.Definition)
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]AppDeploymentFile, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDefinitionSpec.
func (in *AppDefinitionSpec) DeepCopy() *AppDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(AppDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDefinitionStatus) DeepCopyInto(out *AppDefinitionStatus) {
	*out = *in
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDefinitionStatus.
func (in *AppDefinitionStatus) DeepCopy() *AppDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(AppDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentFile) DeepCopyInto(out *AppDeploymentFile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentFile.
func (in *AppDeploymentFile) DeepCopy() *AppDeploymentFile {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AppDeploymentFiles) DeepCopyInto(out *AppDeploymentFiles) {
	{
		in := &in
		*out = make(AppDeploymentFiles, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentFiles.
func (in AppDeploymentFiles) DeepCopy() AppDeploymentFiles {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentFiles)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentJob) DeepCopyInto(out *AppDeploymentJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentJob.
func (in *AppDeploymentJob) DeepCopy() *AppDeploymentJob {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppDeploymentJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentJobComponentStatus) DeepCopyInto(out *AppDeploymentJobComponentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DNSEntries != nil {
		in, out := &in.DNSEntries, &out.DNSEntries
		*out = make([]v1alpha1.DNSEntry, len(*in))
		copy(*out, *in)
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]v1alpha1.Output, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentJobComponentStatus.
func (in *AppDeploymentJobComponentStatus) DeepCopy() *AppDeploymentJobComponentStatus {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentJobComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentJobList) DeepCopyInto(out *AppDeploymentJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppDeploymentJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentJobList.
func (in *AppDeploymentJobList) DeepCopy() *AppDeploymentJobList {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppDeploymentJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentJobSpec) DeepCopyInto(out *AppDeploymentJobSpec) {
	*out = *in
	in.AppDefinition.DeepCopyInto(&out.AppDefinition)
	out.AppEnvRef = in.AppEnvRef
	if in.Only != nil {
		in, out := &in.Only, &out.Only
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make([]v1alpha1.Var, len(*in))
		copy(*out, *in)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make(AppDeploymentFiles, len(*in))
		copy(*out, *in)
	}
	out.CustomTimeout = in.CustomTimeout
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentJobSpec.
func (in *AppDeploymentJobSpec) DeepCopy() *AppDeploymentJobSpec {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentJobStatus) DeepCopyInto(out *AppDeploymentJobStatus) {
	*out = *in
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = new(AppDeploymentPlan)
		(*in).DeepCopyInto(*out)
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]AppDeploymentJobComponentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentJobStatus.
func (in *AppDeploymentJobStatus) DeepCopy() *AppDeploymentJobStatus {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentPlan) DeepCopyInto(out *AppDeploymentPlan) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]AppDeploymentPlanAction, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentPlan.
func (in *AppDeploymentPlan) DeepCopy() *AppDeploymentPlan {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentPlanAction) DeepCopyInto(out *AppDeploymentPlanAction) {
	*out = *in
	out.ResourceKind = in.ResourceKind
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentPlanAction.
func (in *AppDeploymentPlanAction) DeepCopy() *AppDeploymentPlanAction {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentPlanAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentRequest) DeepCopyInto(out *AppDeploymentRequest) {
	*out = *in
	if in.AppDefinitionRef != nil {
		in, out := &in.AppDefinitionRef, &out.AppDefinitionRef
		*out = new(v1alpha1.PlanRef)
		**out = **in
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.CustomTimeout = in.CustomTimeout
	if in.Only != nil {
		in, out := &in.Only, &out.Only
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]AppDeploymentFile, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentRequest.
func (in *AppDeploymentRequest) DeepCopy() *AppDeploymentRequest {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEnv) DeepCopyInto(out *AppEnv) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEnv.
func (in *AppEnv) DeepCopy() *AppEnv {
	if in == nil {
		return nil
	}
	out := new(AppEnv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppEnv) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEnvDeploymentStatus) DeepCopyInto(out *AppEnvDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEnvDeploymentStatus.
func (in *AppEnvDeploymentStatus) DeepCopy() *AppEnvDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(AppEnvDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEnvHostEnv) DeepCopyInto(out *AppEnvHostEnv) {
	*out = *in
	out.CloudAccessConfigRef = in.CloudAccessConfigRef
	out.NamespaceClaimRef = in.NamespaceClaimRef
	out.ClusterRef = in.ClusterRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEnvHostEnv.
func (in *AppEnvHostEnv) DeepCopy() *AppEnvHostEnv {
	if in == nil {
		return nil
	}
	out := new(AppEnvHostEnv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEnvList) DeepCopyInto(out *AppEnvList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppEnv, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEnvList.
func (in *AppEnvList) DeepCopy() *AppEnvList {
	if in == nil {
		return nil
	}
	out := new(AppEnvList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppEnvList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEnvSpec) DeepCopyInto(out *AppEnvSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.CloudAccessConfigRef != nil {
		in, out := &in.CloudAccessConfigRef, &out.CloudAccessConfigRef
		*out = new(v1alpha1.CloudAccessConfigRef)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(int)
		**out = **in
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(AppEnvVars, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEnvSpec.
func (in *AppEnvSpec) DeepCopy() *AppEnvSpec {
	if in == nil {
		return nil
	}
	out := new(AppEnvSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEnvStatus) DeepCopyInto(out *AppEnvStatus) {
	*out = *in
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
	out.HostEnv = in.HostEnv
	if in.CertIssuers != nil {
		in, out := &in.CertIssuers, &out.CertIssuers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IngressClasses != nil {
		in, out := &in.IngressClasses, &out.IngressClasses
		*out = make([]IngressClass, len(*in))
		copy(*out, *in)
	}
	out.Deployment = in.Deployment
	in.DeploymentLastChecked.DeepCopyInto(&out.DeploymentLastChecked)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEnvStatus.
func (in *AppEnvStatus) DeepCopy() *AppEnvStatus {
	if in == nil {
		return nil
	}
	out := new(AppEnvStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEnvVar) DeepCopyInto(out *AppEnvVar) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEnvVar.
func (in *AppEnvVar) DeepCopy() *AppEnvVar {
	if in == nil {
		return nil
	}
	out := new(AppEnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AppEnvVars) DeepCopyInto(out *AppEnvVars) {
	{
		in := &in
		*out = make(AppEnvVars, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEnvVars.
func (in AppEnvVars) DeepCopy() AppEnvVars {
	if in == nil {
		return nil
	}
	out := new(AppEnvVars)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	if in.Deps != nil {
		in, out := &in.Deps, &out.Deps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Plan = in.Plan
	if in.Helm != nil {
		in, out := &in.Helm, &out.Helm
		*out = new(Helm)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentRoles != nil {
		in, out := &in.DeploymentRoles, &out.DeploymentRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make(v1alpha1.InputValues, len(*in))
		copy(*out, *in)
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]OutputDefinition, len(*in))
		copy(*out, *in)
	}
	if in.WorkloadIdentity != nil {
		in, out := &in.WorkloadIdentity, &out.WorkloadIdentity
		*out = new(WorkloadIdentity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDefinition) DeepCopyInto(out *ComponentDefinition) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
	if in.Switch != nil {
		in, out := &in.Switch, &out.Switch
		*out = make([]SwitchableComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDefinition.
func (in *ComponentDefinition) DeepCopy() *ComponentDefinition {
	if in == nil {
		return nil
	}
	out := new(ComponentDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ComponentDefinitions) DeepCopyInto(out *ComponentDefinitions) {
	{
		in := &in
		*out = make(ComponentDefinitions, len(*in))
		for key, val := range *in {
			var outVal *ComponentDefinition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ComponentDefinition)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDefinitions.
func (in ComponentDefinitions) DeepCopy() ComponentDefinitions {
	if in == nil {
		return nil
	}
	out := new(ComponentDefinitions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Helm) DeepCopyInto(out *Helm) {
	*out = *in
	if in.AdditionalValuesFiles != nil {
		in, out := &in.AdditionalValuesFiles, &out.AdditionalValuesFiles
		*out = make([]AdditionalHelmValuesFile, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Helm.
func (in *Helm) DeepCopy() *Helm {
	if in == nil {
		return nil
	}
	out := new(Helm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressClass) DeepCopyInto(out *IngressClass) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressClass.
func (in *IngressClass) DeepCopy() *IngressClass {
	if in == nil {
		return nil
	}
	out := new(IngressClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Input) DeepCopyInto(out *Input) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Input.
func (in *Input) DeepCopy() *Input {
	if in == nil {
		return nil
	}
	out := new(Input)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputDefinition) DeepCopyInto(out *OutputDefinition) {
	*out = *in
	out.OutputDefinition = in.OutputDefinition
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputDefinition.
func (in *OutputDefinition) DeepCopy() *OutputDefinition {
	if in == nil {
		return nil
	}
	out := new(OutputDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretDefinition) DeepCopyInto(out *SecretDefinition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretDefinition.
func (in *SecretDefinition) DeepCopy() *SecretDefinition {
	if in == nil {
		return nil
	}
	out := new(SecretDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in SecretDefinitions) DeepCopyInto(out *SecretDefinitions) {
	{
		in := &in
		*out = make(SecretDefinitions, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretDefinitions.
func (in SecretDefinitions) DeepCopy() SecretDefinitions {
	if in == nil {
		return nil
	}
	out := new(SecretDefinitions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchableComponent) DeepCopyInto(out *SwitchableComponent) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchableComponent.
func (in *SwitchableComponent) DeepCopy() *SwitchableComponent {
	if in == nil {
		return nil
	}
	out := new(SwitchableComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarDefinition) DeepCopyInto(out *VarDefinition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarDefinition.
func (in *VarDefinition) DeepCopy() *VarDefinition {
	if in == nil {
		return nil
	}
	out := new(VarDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in VarDefinitions) DeepCopyInto(out *VarDefinitions) {
	{
		in := &in
		*out = make(VarDefinitions, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarDefinitions.
func (in VarDefinitions) DeepCopy() VarDefinitions {
	if in == nil {
		return nil
	}
	out := new(VarDefinitions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentity) DeepCopyInto(out *WorkloadIdentity) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make([]WorkloadIdentityAccess, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentity.
func (in *WorkloadIdentity) DeepCopy() *WorkloadIdentity {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityAccess) DeepCopyInto(out *WorkloadIdentityAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityAccess.
func (in *WorkloadIdentityAccess) DeepCopy() *WorkloadIdentityAccess {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityAccess)
	in.DeepCopyInto(out)
	return out
}
