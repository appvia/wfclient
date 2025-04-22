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
	"fmt"

	"k8s.io/apimachinery/pkg/types"
)

// Mark 2023-10-27: if only I understood go generics better... fml for this repetition...
// CloudAccessConfigRef is a reference to a cloud access config
type CloudAccessConfigRef struct {
	// Workspace identifies the workspace in which this item exists. If empty, assume the same
	// workspace as the object holding the reference.
	Workspace WorkspaceKey `json:"workspace,omitempty"`
	// Name is the object name of the referenced item
	Name string `json:"name"`
}

// Set sets this reference to point to the provided object
func (w *CloudAccessConfigRef) Set(o Object) {
	w.Name = o.GetName()
	w.Workspace = Workspace(o)
}

func (w *CloudAccessConfigRef) Is(o Object) bool {
	if w == nil || o == nil {
		return w == nil && o == nil
	}
	return o.GetName() == w.Name && Workspace(o) == w.Workspace
}

func (w *CloudAccessConfigRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      w.Name,
		Namespace: w.Workspace.Namespace(),
	}
}

func (w *CloudAccessConfigRef) Equals(other CloudAccessConfigRef) bool {
	return w.Workspace == other.Workspace && w.Name == other.Name
}

// Empty will return true if the reference is nil or the name is unpopulated (it does not check the
// workspace as that is optional and unusable if no name is provided anyway).
func (w *CloudAccessConfigRef) Empty() bool {
	return w == nil || w.Name == ""
}

// ClusterRef is a reference to a cluster
type ClusterRef struct {
	// Workspace identifies the workspace in which this item exists. If empty, assume the same
	// workspace as the object holding the reference.
	Workspace WorkspaceKey `json:"workspace,omitempty"`
	// Name is the object name of the referenced item
	Name string `json:"name"`
}

// Set sets this reference to point to the provided object
func (w *ClusterRef) Set(o Object) {
	w.Name = o.GetName()
	w.Workspace = Workspace(o)
}

func (w *ClusterRef) Is(o Object) bool {
	if w == nil || o == nil {
		return w == nil && o == nil
	}
	return o.GetName() == w.Name && Workspace(o) == w.Workspace
}

func (w *ClusterRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      w.Name,
		Namespace: w.Workspace.Namespace(),
	}
}

func (w *ClusterRef) Equals(other ClusterRef) bool {
	return w.Workspace == other.Workspace && w.Name == other.Name
}

func (w *ClusterRef) Empty() bool {
	if w == nil {
		return true
	}
	return w.Workspace == "" && w.Name == ""
}

// NamespaceClaimRef is a reference to a namespace claim
type NamespaceClaimRef struct {
	// Workspace identifies the workspace in which this item exists. If empty, assume the same
	// workspace as the object holding the reference.
	Workspace WorkspaceKey `json:"workspace,omitempty"`
	// Name is the object name of the referenced item
	Name string `json:"name"`
}

// Set sets this reference to point to the provided object
func (w *NamespaceClaimRef) Set(o Object) {
	w.Name = o.GetName()
	w.Workspace = Workspace(o)
}

func (w *NamespaceClaimRef) Is(o Object) bool {
	if w == nil || o == nil {
		return w == nil && o == nil
	}
	return o.GetName() == w.Name && Workspace(o) == w.Workspace
}

func (w *NamespaceClaimRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      w.Name,
		Namespace: w.Workspace.Namespace(),
	}
}

func (w *NamespaceClaimRef) Equals(other NamespaceClaimRef) bool {
	return w.Workspace == other.Workspace && w.Name == other.Name
}

func (w *NamespaceClaimRef) Empty() bool {
	return w.Workspace == "" && w.Name == ""
}

// AppEnvRef is a reference to an appenv
type AppEnvRef struct {
	// Workspace identifies the workspace in which this item exists. If empty, assume the same
	// workspace as the object holding the reference.
	Workspace WorkspaceKey `json:"workspace,omitempty"`
	// App is the spec.application of the appenv
	App string `json:"app"`
	// EnvName is the spec.name of the appenv
	EnvName string `json:"envName"`
}

func (w *AppEnvRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      w.AppEnvObjectName(),
		Namespace: w.Workspace.Namespace(),
	}
}

func (w *AppEnvRef) AppEnvObjectName() string {
	if w == nil {
		return ""
	}
	if w.App == "" || w.EnvName == "" {
		return ""
	}
	// Assume appenvs are named APP-ENV (this is validated by server/apps/handlers/internal/appenv/validate.go)
	return fmt.Sprintf("%s-%s", w.App, w.EnvName)
}

func (w *AppEnvRef) Equals(other AppEnvRef) bool {
	return w.Workspace == other.Workspace && w.App == other.App && w.EnvName == other.EnvName
}

func (w *AppEnvRef) Empty() bool {
	return w.Workspace == "" && w.App == "" && w.EnvName == ""
}

// DNSZoneRef is a reference to a global or workspaced DNS zone
type DNSZoneRef struct {
	// Workspace identifies the workspace in which this item exists. If empty, this refers to a
	// GlobalDNSZone, if populated, a workspaced DNSZone
	Workspace WorkspaceKey `json:"workspace,omitempty"`
	// Name is the object name of the referenced item
	Name string `json:"name"`
}

// Set sets this reference to point to the provided object
func (w *DNSZoneRef) Set(o Object) {
	w.Name = o.GetName()
	w.Workspace = Workspace(o)
}

func (w *DNSZoneRef) Is(o Object) bool {
	if w == nil || o == nil {
		return w == nil && o == nil
	}
	return o.GetName() == w.Name && Workspace(o) == w.Workspace
}

func (w *DNSZoneRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      w.Name,
		Namespace: w.Workspace.Namespace(),
	}
}

func (w *DNSZoneRef) Equals(other DNSZoneRef) bool {
	return w.Workspace == other.Workspace && w.Name == other.Name
}

func (w *DNSZoneRef) Empty() bool {
	return w.Workspace == "" && w.Name == ""
}

// PlatformSecretRef is a reference to a platform secret
type PlatformSecretRef string

func (p PlatformSecretRef) String() string {
	return string(p)
}

type PackageRef struct {
	Name    string        `json:"name"`
	Version ObjectVersion `json:"version"`
}

// ToObjectName returns the underlying (internal, k8s) object name for the version of the package
func (w *PackageRef) ToObjectName() string {
	return w.Version.ToVersionedName(w.Name)
}

func (w *PackageRef) IsTo(o Object) bool {
	if w == nil || o == nil {
		return w == nil && o == nil
	}
	if oVer, ok := o.(VersionedObject); ok {
		return w.Name == oVer.VersionOf() && w.Version == oVer.GetVersion()
	}

	return false
}

func (w *PackageRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name: w.Version.ToVersionedName(w.Name),
	}
}

func (w *PackageRef) Equals(other PackageRef) bool {
	return w.Version == other.Version && w.Name == other.Name
}

func (w *PackageRef) Empty() bool {
	return w.Version == "" && w.Name == ""
}

// PlanRef represents a versioned reference to a plan construct
type PlanRef struct {
	Name            string        `json:"name"`
	Version         ObjectVersion `json:"version"`
	FollowPublished bool          `json:"followPublished,omitempty"`
}

// ToObjectName returns the underlying (internal, k8s) object name for the version of the plan
func (w *PlanRef) ToObjectName() string {
	if w == nil {
		return ""
	}
	return w.Version.ToVersionedName(w.Name)
}

func (w *PlanRef) String() string {
	if w == nil {
		return ""
	}
	return w.Name + " (" + string(w.Version) + ")"
}

func (w PlanRef) IsTo(o Object) bool {
	if o == nil {
		return false
	}
	if oVer, ok := o.(VersionedObject); ok {
		return w.Name == oVer.VersionOf() && w.Version == oVer.GetVersion()
	}

	return false
}

func (w *PlanRef) Equals(other *PlanRef) bool {
	if w == nil || other == nil {
		return w == nil && other == nil
	}
	return w.Version == other.Version && w.Name == other.Name
}

func (w PlanRef) Empty() bool {
	return w.Version == "" && w.Name == ""
}

// ClusterNetworkRef is a reference to a cluster network
type ClusterNetworkRef struct {
	// Workspace identifies the workspace in which this item exists. If empty, assume the same
	// workspace as the object holding the reference.
	Workspace WorkspaceKey `json:"workspace,omitempty"`
	// Name is the object name of the referenced item
	Name string `json:"name"`
}

// Set sets this reference to point to the provided object
func (w *ClusterNetworkRef) Set(o Object) {
	w.Name = o.GetName()
	w.Workspace = Workspace(o)
}

func (w *ClusterNetworkRef) Is(o Object) bool {
	if w == nil || o == nil {
		return w == nil && o == nil
	}
	return o.GetName() == w.Name && Workspace(o) == w.Workspace
}

func (w *ClusterNetworkRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      w.Name,
		Namespace: w.Workspace.Namespace(),
	}
}

func (w *ClusterNetworkRef) Equals(other CloudAccessConfigRef) bool {
	return w.Workspace == other.Workspace && w.Name == other.Name
}

// Empty will return true if the reference is nil or the name is unpopulated (it does not check the
// workspace as that is optional and unusable if no name is provided anyway).
func (w *ClusterNetworkRef) Empty() bool {
	return w == nil || w.Name == ""
}

type CloudResourceRef struct {
	// Name is the object name of the referenced item
	Name string `json:"name"`
}

// Set sets this reference to point to the provided object
func (w *CloudResourceRef) Set(o Object) {
	w.Name = o.GetName()
}

func (w *CloudResourceRef) Is(o Object) bool {
	if w == nil || o == nil {
		return w == nil && o == nil
	}
	return o.GetName() == w.Name
}

func (w *CloudResourceRef) Equals(other CloudAccessConfigRef) bool {
	return w.Name == other.Name
}

// Empty will return true if the reference is nil or the name is unpopulated
func (w *CloudResourceRef) Empty() bool {
	return w == nil || w.Name == ""
}
