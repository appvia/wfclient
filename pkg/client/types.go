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
	"context"
	"fmt"
	"io"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
	"github.com/appvia/wfclient/pkg/client/config"
	"github.com/appvia/wfclient/pkg/utils/validation"
)

// ObjectKey identifies a Wayfinder object
type ObjectKey struct {
	Workspace corev1.WorkspaceKey
	Name      string
	Version   corev1.ObjectVersion
}

func (o ObjectKey) String() string {
	if o.Workspace.Key() != "" {
		return fmt.Sprintf("%s/%s", o.Workspace.Key(), o.Name)
	}
	return o.Name
}

// ObjectKeyFromObject returns the ObjectKey given a runtime.Object.
func ObjectKeyFromObject(obj Object) ObjectKey {
	key := ObjectKey{Workspace: corev1.Workspace(obj), Name: obj.GetName()}
	if corev1.IsVersioned(obj) {
		key.Version = obj.(corev1.Versioned).GetVersion()
	}
	return key
}

// Object is the basic interface implemented by all Wayfinder objects
type Object corev1.Object

// ObjectList is the basic list interface implemented by all Wayfinder list objects
type ObjectList corev1.ObjectList

// Reader knows how to read and list Wayfinder objects.
type Reader interface {
	// Get retrieves an obj for the given object key from the Kubernetes Cluster.
	// obj must be a struct pointer so that obj can be updated with the response
	// returned by the Server.
	Get(ctx context.Context, key ObjectKey, obj Object) error

	// List retrieves list of objects for a given namespace and list options. On a
	// successful call, Items field in the list will be populated with the
	// result returned from the server.
	List(ctx context.Context, list ObjectList, opts ...ListOption) error

	// ListVersions lists the available versions of the named versioned object.
	ListVersions(ctx context.Context, name string, list ObjectList, opts ...ListOption) error
}

// Writer knows how to create, delete, and update Kubernetes objects.
type Writer interface {
	// Create saves the object obj in the Wayfinder API.
	Create(ctx context.Context, obj Object, opts ...CreateOption) error

	// Delete deletes the given obj from Wayfinder API.
	Delete(ctx context.Context, obj Object, opts ...DeleteOption) error

	// DeleteAllVersions deletes all versions of the named versioned object from Wayfinder API. Do
	// not specify key.Version when using this method.
	DeleteAllVersions(tx context.Context, key ObjectKey, list ObjectList, opts ...DeleteOption) error

	// Update updates the given obj in the Wayfinder API. obj must be a
	// struct pointer so that obj can be updated with the content returned by the Server.
	Update(ctx context.Context, obj Object, opts ...UpdateOption) error
}

// WFClient provides a simple client to Wayfinder, inspired by controller-manager's client.Client
type WFClient interface {
	Reader
	Writer
	// EndpointRequest allows access to the raw underlying request interface for
	// non-resource API operations
	EndpointRequest(ctx context.Context, endpoint string) RestInterface
	// ResourceRequest allows access to the raw underlying request interface for resource-oriented
	// operations against the type represented by resObj, such as accessing arbitrary sub-resources
	ResourceRequest(ctx context.Context, resObj corev1.Object) RestInterface
	// ResourceClient retrieves the underlying resource-oriented client for this WFClient.
	ResourceClient() Interface
}

// Interface is the api client interface
type Interface interface {
	// Request creates a request instance
	Request() RestInterface
	// Config returns the underlying configuration for the client
	Config() *config.Config
	// CurrentProfile returns the current profile
	CurrentProfile() string
	// OverrideProfile allows you set the selected profile
	OverrideProfile(string) Interface
	// RefreshIdentity is used to refresh the identity token of the user
	RefreshIdentity() error
	// CheckServer ensures any initialization of the selected server profile is done. If saveProfile
	// is true, the client will persist any changes to the profile. If force is true, it will always
	// ping the server, if false, it will only do that if the selected profile does not have API
	// info already set in it
	CheckServer(force, saveProfile bool) error
}

// UpdateHandlerFunc is external method when the configuration has been updated
type UpdateHandlerFunc func() error

// ClientVersionHeader is the header used in the client cli
const ClientVersionHeader = "X-Client-Version"
const ObjectModifiedError = "the object has been modified, please try again"

// RestInterface provides the rest interface
type RestInterface interface {
	// Authorization allows you to override the authorization
	Authorization(string) RestInterface
	// Body returns the body if any
	Body() io.Reader
	// Context sets the request context
	Context(context.Context) RestInterface
	// Create performs a post request
	Create() RestInterface
	// Delete performs a delete
	Delete() RestInterface
	// Do returns the response and error
	Do() (RestInterface, error)
	// Duplicate copes the current request
	Duplicate() RestInterface
	// Endpoint defines the endpoint to use
	Endpoint(string) RestInterface
	// RawEndpoint defines the endpoint to use without a prefix of the server base path
	RawEndpoint(string) RestInterface
	// Exists checks if the resource exists
	Exists() (bool, error)
	// Error returns the error if any
	Error() error
	// Follow indicates we follow the stream from the server if any
	Follow(bool) RestInterface
	// Get performs a get request
	Get() RestInterface
	// GetPayload returns the payload for inspection
	GetPayload() interface{}
	// Warnings returns the headers warnings, if any
	GetWarnings() []validation.Warning
	// WithWarningHandler defines how the client handles response warnings (default is to output them to stderr)
	WithWarningHandler(WarningHandler) RestInterface
	// HasParameter checks if the parameter is set
	HasParameter(string) (string, bool)
	// Name sets the resource name
	Name(string) RestInterface
	// Resource sets the resource and version in the request from the provided source. This is a
	// shortcut for calling ResourceGroup(), ResoucreVersion() and ResourceKind()
	Resource(VersionedResourceSource) RestInterface
	// ResourceAPIVersion sets the API version for the resource to the specified value. Use to
	// override the default API version set by a call to Resource().
	// This will no-op if an empty string is provided, keeping the version set by a call to
	// Resource(), so can safely be called with a potentially empty parameter.
	ResourceAPIVersion(v string) RestInterface
	// ResourceVersion sets the version of a versioned resource we are operating with
	ResourceVersion(rv string) RestInterface

	// Parameters defines a list of parameters for the request
	Parameters(...ParameterFunc) RestInterface
	// Payload set the payload of the request
	Payload(interface{}) RestInterface
	// Post performs a post request
	Post() RestInterface
	// Result set the object which we should decode into
	Result(interface{}) RestInterface
	// Stream is used to stream from the api server
	// Stream(context.Context) (chan []byte, error)
	// SubResource adds a subresource to the operation
	SubResource(string) RestInterface
	// SubResourceName sets the sub-resource name
	SubResourceName(string) RestInterface
	// Workspace set the workspace
	Workspace(corev1.WorkspaceKey) RestInterface
	// Unauthenticated indicates no need to add auth
	Unauthenticated() RestInterface
	// Update performs an put request
	Update() RestInterface
}

// VersionedResourceSource is an object that can describe the version and API name of a resource
type VersionedResourceSource interface {
	GetAPIName() string
	// GetGroupVersion should return the group/version info for the resource
	GetGroupVersion() metav1.GroupVersion
	// IsResourceVersioned should return true if this resource uses resource versioning at the
	// specified API version
	IsResourceVersioned(ver string) bool
}

// ParameterFunc defines a method for a parameter type
type ParameterFunc func() (Parameter, error)

// Parameter is a param to the raw endpoint
type Parameter struct {
	// IsPath indicates if it's a path or query parameter
	IsPath bool
	// Name is the name of the parameter
	Name string
	// Value is the value of the parameter
	Value string
}

// APIError is the client-side representation of an error returned by this client where a structued
// API error is returned from the server
type APIError struct {
	// Code is an optional machine readable code used to describe the code
	Code int `json:"code"`
	// Detail is the actual error thrown by the upstream
	Detail string `json:"detail"`
	// Message is a human readable message related to the error
	Message string `json:"message"`
	// URI is the uri of the request
	URI string `json:"uri"`
	// Verb was the http request verb used
	Verb string `json:"verb"`
	// Validation will be populated with the underlying structured validation error if applicable
	Validation *validation.Error
	// DependencyViolation will be populated with the underlying structured dependency violation
	// error if applicable
	DependencyViolation *validation.ErrDependencyViolation
}

// Error returns the error message
func (e APIError) Error() string {
	return e.Message
}

// Is reports whether the error message matches the target error message
func (e APIError) Is(target error) bool {
	return e.Message == target.Error()
}

type WarningHandler func(context.Context, []validation.Warning)

type RequestDo func(req *http.Request) (*http.Response, error)
