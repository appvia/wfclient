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
	// WFAnnotationPrefix is the prefix used for all Wayfinder annotations
	WFAnnotationPrefix = "appvia.io"
	// VersioningAnnotationName is the name used for versioning annotations
	VersioningAnnotationName = "versioning"
)

// WFAnnotation returns an annotation in the format 'appvia.io/tag'
func WFAnnotation(tag string) string {
	return fmt.Sprintf("%s/%s", WFAnnotationPrefix, tag)
}

// WFAnnotationNamed returns a named annotation in the format 'name.appvia.io/tag'
func WFAnnotationNamed(name, tag string) string {
	return fmt.Sprintf("%s.%s/%s", name, WFAnnotationPrefix, tag)
}

func VersioningAnnotation(field string) string {
	return WFAnnotationNamed(VersioningAnnotationName, field)
}

var (
	// AnnotationCloudResourceName defines the name in cloud of a resource, used to ensure that the
	// name is persisted even if we change our naming convention in future.
	AnnotationCloudResourceName = WFAnnotation("cloudResourceName")

	// AnnotationPackageProviderFilter is used to further filter packages with no workload identity
	// This allows us to restrict very specific packages to a very specific cluster provider
	AnnotationPackageProviderFilter = WFAnnotation("packageProviderFilter")

	// AnnotationPublished indicates whether this resource is in the 'published' state. It should be
	// the string 'true' or 'false'.
	AnnotationPublished = WFAnnotation("published")

	// AnnotationPublishState is an optional annotation to indicate an informative reason why the
	// resource is published or not published, e.g. because it is in draft etc.
	AnnotationPublishState = WFAnnotation("publishState")

	// AnnotationRefresh is used to force re-reconciliation of a resource. If this annotation is
	// present and its value is different to the value status.LastReconcileRefresh() then the
	// resource will be re-reconciled.
	AnnotationRefresh = WFAnnotation("refresh")

	// VersioningAnnotationLegacyPlan is set on legacy unversioned objects of resources that are now
	// versioned, when viewed at the later, versioned, API version.
	VersioningAnnotationLegacyPlan = VersioningAnnotation("legacyPlan")
)

const (
	AnnotationValuePublishedTrue  = "true"
	AnnotationValuePublishedFalse = "false"
)

func IsPublished(obj Object) bool {
	return obj.GetAnnotations()[AnnotationPublished] == AnnotationValuePublishedTrue || obj.GetAnnotations()[AnnotationPublished] == ""
}

func SetRefresh(object Object, refresh string) {
	if object.GetAnnotations() == nil {
		object.SetAnnotations(map[string]string{})
	}
	object.GetAnnotations()[AnnotationRefresh] = refresh
}

// IsRefreshAnnotationUpToDate will return true if this object's refresh annotation matches the provided
// refresh string (or the provided refresh string is empty)
func IsRefreshAnnotationUpToDate(object Object, refresh string) bool {
	if refresh == "" {
		return true
	}
	if object.GetAnnotations()[AnnotationRefresh] == refresh {
		return true
	}
	return false
}

// IsRefreshProcessed will return true if the current refresh value (in annotation
// appvia.io/refresh) has been processed by the relevant controller (i.e. the status field
// status.lastReconcile.refresh has a value equal to the annotation appvia.io/refresh)
func IsRefreshProcessed(object Object) bool {
	return object.GetAnnotations()[AnnotationRefresh] == "" || object.GetAnnotations()[AnnotationRefresh] == object.GetCommonStatus().LastReconcileRefresh()
}

// IsReconciled returns true if the generation has been considered by the controller and there is no
// pending refresh.
func IsReconciled(object Object) bool {
	return object.GetGeneration() == object.GetCommonStatus().LastReconcileGeneration() && IsRefreshProcessed(object)
}
