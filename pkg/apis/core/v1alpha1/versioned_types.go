/**
 * Copyright 2025 Appvia Ltd <info@appvia.io>
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
	"strings"
)

// ObjectVersion identifies the version of a versionable resource.
//
// Must be a valid semver in the format X.Y.Z. An optional 'v' prefix is supported but discouraged.
// An optional -suffix can be provided; note in semver that 1.0.0 is _after_ 1.0.0-1.
// +kubebuilder:validation:Pattern=`^(v|)([0-9]+\.[0-9]+\.[0-9]+(-[a-z0-9-.]+|)|)$`
type ObjectVersion string

func (o ObjectVersion) ToVersionedName(name string) string {
	if o == "" {
		return name
	}
	return fmt.Sprintf("%s.%s", name, o)
}

func (o ObjectVersion) String() string {
	return string(o)
}

// Versioned is the interface implemented by types which are resource versioned
// +kubebuilder:object:generate=false
type Versioned interface {
	// VersionOf returns the name that this is a version of
	VersionOf() string
	// GetVersion returns the object version of this object
	GetVersion() ObjectVersion
	// SetVersion sets the current version of this versioned object
	SetVersion(ObjectVersion)
	// SetTags sets the tags of this versioned object
	SetTags([]string)
	// SetDescription sets the description of the versioned object
	SetDescription(string)
}

// PlanMetadata describes the metadata of a plan
// +kubebuilder:object:generate=false
type PlanMetadata struct {
	Tags        []string `json:"tags,omitempty"`
	Description string   `json:"description,omitempty"`
}

func SetObjectTags(obj Object, tags []string) {
	labels := map[string]string{}
	label := "category.appvia.io/"

	// all labels not including tags
	for k, v := range obj.GetLabels() {
		if !strings.HasPrefix(k, label) {
			labels[k] = v
		}
	}

	for _, tag := range tags {
		kstring := strings.ToLower(strings.ReplaceAll(tag, " ", ""))
		labels[label+kstring] = tag
	}

	obj.SetLabels(labels)
}

// +kubebuilder:object:generate=false
type VersionedObject interface {
	Object
	Versioned
}

// IsVersioned returns true if the provided object supports resource versioning
func IsVersioned(obj Object) bool {
	_, ok := obj.(Versioned)
	return ok
}

// GetVersion returns the version of the provided object, or an empty string if the object is not
// versioned
func GetVersion(obj Object) ObjectVersion {
	if versioned, ok := obj.(Versioned); ok {
		return versioned.GetVersion()
	}
	return ""
}

const QueryParamAllVersions = "allVersions"
