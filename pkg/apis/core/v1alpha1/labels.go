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

import "fmt"

const (
	// WFLabelPrefix is the prefix used for all Wayfinder labels
	WFLabelPrefix = "appvia.io"
)

// Label returns a wayfinder label on a resource
func Label(tag string) string {
	return fmt.Sprintf("%s/%s", WFLabelPrefix, tag)
}

// NamedLabel returns a wayfinder named label on a resource
func NamedLabel(name, tag string) string {
	return fmt.Sprintf("%s.%s/%s", name, WFLabelPrefix, tag)
}

var (
	// LabelVersionOf is the name that a versioned resource is a version of
	LabelVersionOf = Label("versionOf")

	// LabelZoneType is the type of zone
	LabelZoneType = Label("zoneType")
)
