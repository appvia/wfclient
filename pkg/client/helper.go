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

package client

import (
	"net/http"
	"strings"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
)

// IsNotFound check if error is an 404 error
func IsNotFound(err error) bool {
	return isExpectedError(err, http.StatusNotFound)
}

// IsNotAuthorized checks if the error as a 401
func IsNotAuthorized(err error) bool {
	return isExpectedError(err, http.StatusUnauthorized)
}

// IsNotImplemented check if the error is a 501
func IsNotImplemented(err error) bool {
	return isExpectedError(err, http.StatusNotImplemented)
}

// IsNotAllowed checks if the response was a 403 forbidden
func IsNotAllowed(err error) bool {
	return isExpectedError(err, http.StatusForbidden)
}

// IsServiceUnavailable checks for s 503
func IsServiceUnavailable(err error) bool {
	return isExpectedError(err, http.StatusServiceUnavailable)
}

// IsMethodNotAllowed checks if the response was a 405 forbidden
func IsMethodNotAllowed(err error) bool {
	return isExpectedError(err, http.StatusMethodNotAllowed)
}

// IsBadRequest checks if the response was a 400
func IsBadRequest(err error) bool {
	return isExpectedError(err, http.StatusBadRequest)
}

// isExpectError checks if the error an apiError and compares the code
func isExpectedError(err error, code int) bool {
	e, ok := (err).(*APIError)
	if !ok {
		return false
	}

	return e.Code == code
}

func IsAlreadyExists(err error) bool {
	return kerrors.IsAlreadyExists(err) ||
		strings.Contains(err.Error(), "already exists") // TODO: Fix error types we return
}

func IsObjectModified(err error) bool {
	return err != nil && err.Error() == ObjectModifiedError
}

// For returns a versioned resource source for the provided object
func For(obj corev1.Object) VersionedResourceSource {
	return resSrc{obj}
}

// resSrc is a wrapper for corev1.Object that implements VersionedResourceSource
type resSrc struct {
	obj corev1.Object
}

// GetGroupVersion returns the API version for this resource
func (v resSrc) GetGroupVersion() metav1.GroupVersion {
	gvk := v.obj.GetObjectKind().GroupVersionKind()

	// // If we've been passed a non-initialized object, it's entirely possible that these are
	// // unpopulated, in which case, look them up from the schema.
	// if gvk.Group == "" || gvk.Version == "" {
	// 	// This might not find the object, but even if it does, let's continue here - the errors
	// 	// will come out more meaningfully when we try and use this against the API.
	// 	gvk, _ = schema.GetGroupKindVersion(v.obj)
	// }

	return metav1.GroupVersion{
		Group:   gvk.Group,
		Version: gvk.Version,
	}
}

// GetAPIName returns the API name for this resource
func (v resSrc) GetAPIName() string {
	return v.obj.APIPath()
}

func (v resSrc) IsResourceVersioned(ver string) bool {
	return corev1.IsVersioned(v.obj)
}
