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

// WayfinderResource is reference to a resource within wayfinder
// +k8s:openapi-gen=true
type WayfinderResource struct {
	// Kind is the kind of resource
	// +kubebuilder:validation:Optional
	Kind string `json:"kind,omitempty"`
	// Name is the name of the resource
	// +kubebuilder:validation:Optional
	Name string `json:"name,omitempty"`
	// Workspace is the workspace of the resource
	// +kubebuilder:validation:Optional
	Workspace WorkspaceKey `json:"workspace,omitempty"`
}

// CloudResource is a reference to a resource in the given cloud
// +k8s:openapi-gen=true
type CloudResource struct {
	// Name is the name of the resource in the cloud
	// +kubebuilder:validation:Optional
	Name string `json:"name,omitempty"`
	// ID is the full cloud ID of the resource, e.g. the full ARN, Azure Resource ID, etc.
	// +kubebuilder:validation:Optional
	ID string `json:"id,omitempty"`
	// Type is the type of resource e.g. AzureManagedIdentity, AWSIAMRole
	// +kubebuilder:validation:Optional
	Type string `json:"type,omitempty"`
	// Cloud is the cloud where the resource exists
	// +kubebuilder:validation:Optional
	Cloud string `json:"cloud,omitempty"`
}

// OwnedResources is a reference to the resources created and owned by an object in Wayfinder
// +k8s:openapi-gen=true
type OwnedResources struct {
	// WayfinderResources are resources within wayfinder
	// +kubebuilder:validation:Type=array
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +kubebuilder:validation:Optional
	WayfinderResources []WayfinderResource `json:"wayfinderResources,omitempty"`
	// CloudResources are managed resources in the relevant cloud
	// +kubebuilder:validation:Type=array
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +kubebuilder:validation:Optional
	CloudResources []CloudResource `json:"cloudResources,omitempty"`
}

func (or *OwnedResources) ContainsCloudResource(name, cloud, rtype string) bool {
	if or == nil {
		return false
	}
	for _, res := range or.CloudResources {
		if res.Name == name && res.Cloud == cloud && res.Type == rtype {
			return true
		}
	}

	return false
}

func (or *OwnedResources) GetFirstCloudResourceOfType(rtype string) *CloudResource {
	if or == nil {
		return nil
	}
	for _, res := range or.CloudResources {
		if res.Type == rtype {
			return &res
		}
	}

	return nil
}

func (or *OwnedResources) ContainsWayfinderResource(kind, name string, workspace WorkspaceKey) bool {
	if or == nil {
		return false
	}
	for _, res := range or.WayfinderResources {
		if res.Kind == kind && res.Name == name && (res.Workspace == "" || res.Workspace == workspace) {
			return true
		}
	}
	return false
}

func (or *OwnedResources) AddWayfinderResource(res WayfinderResource) OwnedResources {
	if or.ContainsWayfinderResource(res.Kind, res.Name, res.Workspace) {
		return *or
	}

	or.WayfinderResources = append(or.WayfinderResources, res)

	return *or
}

func (or *OwnedResources) RemoveWayfinderResource(res WayfinderResource) OwnedResources {
	if !or.ContainsWayfinderResource(res.Kind, res.Name, res.Workspace) {
		return *or
	}
	newRes := []WayfinderResource{}
	for i, r := range or.WayfinderResources {
		if r.Kind != res.Kind || r.Name != res.Name || r.Workspace != res.Workspace {
			newRes = append(newRes, or.WayfinderResources[i])
		}
	}
	or.WayfinderResources = newRes

	return *or
}

func (or *OwnedResources) AddCloudResource(res CloudResource) OwnedResources {
	if or.ContainsCloudResource(res.Name, res.Cloud, res.Type) {
		return *or
	}

	or.CloudResources = append(or.CloudResources, res)

	return *or
}

func (or *OwnedResources) RemoveCloudResource(typ, name string) OwnedResources {
	newRes := []CloudResource{}
	for i, r := range or.CloudResources {
		if r.Type != typ || r.Name != name {
			newRes = append(newRes, or.CloudResources[i])
		}
	}
	or.CloudResources = newRes

	return *or
}
