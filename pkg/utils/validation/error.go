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

package validation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
	"github.com/appvia/wfclient/pkg/common"
)

// FieldRoot is used to reference the root object
const FieldRoot = "(root)"

// ErrNonSystemServiceAccount indicates that the resource can only be managed by Wayfinder itself
var ErrNonSystemServiceAccount = NewError("operations on this resource can only be performed by Wayfinder")

// Error is a specific error returned when the input provided by
// the user has failed validation somehow.
type Error struct {
	// Code is an optional machine readable code used to describe the code
	Code int `json:"code"`
	// Message is a human readable message related to the error
	Message string `json:"message"`
	// FieldErrors are the individual validation errors found against the submitted data.
	FieldErrors []FieldError `json:"fieldErrors"`
}

// Gets all the field errors which are warnings
func (e Error) GetWarnings() []FieldError {
	warnings := []FieldError{}
	for _, fe := range e.FieldErrors {
		switch fe.ErrCode {
		case Deprecated, ShouldExist, FieldWarning:
			warnings = append(warnings, fe)
		}
	}
	return warnings
}

// Gets all the field errors which are NOT warnings
func (e Error) GetNonWarnings() []FieldError {
	nonWarnings := []FieldError{}
	for _, fe := range e.FieldErrors {
		switch fe.ErrCode {
		case Deprecated, ShouldExist, FieldWarning:
		default:
			nonWarnings = append(nonWarnings, fe)
		}
	}
	return nonWarnings
}

// NewError returns a new validation error.
func NewError(msg string) *Error {
	return &Error{
		Code:    400,
		Message: msg,
	}
}

// NewErrorf returns a new validation error, with formatting.
func NewErrorf(format string, args ...interface{}) *Error {
	return &Error{
		Code:    400,
		Message: fmt.Sprintf(strings.TrimRight(format, ":\n"), args...),
	}
}

func NewFieldError(field string, errCode ErrorCode, message string) FieldError {
	return FieldError{
		Field:   field,
		ErrCode: errCode,
		Message: message,
	}
}

func NewFieldErrorf(field string, errCode ErrorCode, format string, args ...interface{}) FieldError {
	return FieldError{
		Field:   field,
		ErrCode: errCode,
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns the details of the validation error.
func (e Error) Error() string {
	if len(e.FieldErrors) == 0 {
		return ""
	}
	sb := strings.Builder{}
	sb.WriteString(e.Message)
	sb.WriteString(":\n")
	for _, fe := range e.FieldErrors {
		if fe.Field == FieldRoot {
			sb.WriteString(fmt.Sprintf(" * %s\n", fe.Message))
		} else {
			sb.WriteString(fmt.Sprintf(" * %s: %s\n", fe.Field, fe.Message))
		}
	}
	return sb.String()
}

// HasErrors returns true if any field errors have been added to this validation error.
func (e *Error) HasErrors() bool {
	return len(e.FieldErrors) > 0
}

func (e *Error) HasErrorContaining(needle string) bool {
	for _, x := range e.FieldErrors {
		if strings.Contains(x.Message, needle) {
			return true
		}
	}

	return false
}

// HasValidationError checks if the validation error is present
func (e *Error) HasValidationError(v FieldError) bool {
	for _, x := range e.FieldErrors {
		switch {
		case x.Field != v.Field:
			fallthrough
		case x.ErrCode != v.ErrCode:
			fallthrough
		case x.Message != v.Message:
			continue
		}

		return true
	}

	return false
}

// AddNewFieldErrors adds multiple strongly typed field errors without duplication
func (e *Error) AddNewFieldErrors(fes []FieldError) {
	// Only add non duplicates
	for _, fe := range fes {
		e.AddNewFieldError(fe)
	}
}

// AddNewFieldError adds a strongly typed field errors without duplication
func (e *Error) AddNewFieldError(fe FieldError) {
	if !e.IsPresent(fe) {
		e.FieldErrors = append(e.FieldErrors, fe)
	}
}

// WithFieldError adds a field error to the validation error and returns it for fluent loveliness.
func (e *Error) WithFieldError(field string, errCode ErrorCode, message string) *Error {
	return e.AddFieldError(field, errCode, message)
}

// WithFieldErrorf adds an error for a specific field to a validation error.
func (e *Error) WithFieldErrorf(field string, errCode ErrorCode, format string, args ...interface{}) *Error {
	return e.AddFieldErrorf(field, errCode, format, args...)
}

// IsPresent will determin if this filed error has already been added
func (e *Error) IsPresent(fe FieldError) bool {
	for _, e := range e.FieldErrors {
		if e == fe {
			return true
		}
	}

	return false
}

// AddFieldError adds an error for a specific field to a validation error.
func (e *Error) AddFieldError(field string, errCode ErrorCode, message string) *Error {
	// use strongly typed
	e.AddNewFieldErrors([]FieldError{NewFieldError(field, errCode, message)})

	return e
}

// AddFieldErrorf adds a general (invalid value) validation warning error for a specific field
func (e *Error) AddFieldWarningf(ctx context.Context, field string, message string, args ...interface{}) {
	warning := Warning{
		WarningType: WarningTypeGeneral,
		Message:     field + ": " + fmt.Sprintf(message, args...),
	}

	e.AddFieldError(field, FieldWarning, warningMessage(ctx, warning))
}

// AddFieldDependencyWarning adds a validation warning as an error
func (e *Error) AddFieldDependencyWarning(ctx context.Context, field string, name, kind string, workspace corev1.WorkspaceKey) {
	warning := Warning{
		WarningType: WarningTypeDependency,
		Kind:        kind,
		Name:        name,
		Workspace:   workspace,
	}

	e.AddFieldError(field, ShouldExist, warningMessage(ctx, warning))
}

// AddDependencyWarningOrError adds either a warning (on dry-run) or a MustExist error for the
// missing dependency
func (e *Error) AddDependencyWarningOrError(ctx context.Context, dryRun bool, field string, name, kind string, workspace corev1.WorkspaceKey) {
	if dryRun {
		e.AddFieldDependencyWarning(ctx, field, name, kind, workspace)
	} else {
		e.AddFieldErrorf(field, MustExist, "%s %s does not exist", kind, name)
	}
}

// AddVersionedDependencyWarningOrError adds either a warning (on dry-run) or a MustExist error for the
// missing dependency
func (e *Error) AddVersionedDependencyWarningOrError(ctx context.Context, dryRun bool, field string, name, kind string, version corev1.ObjectVersion, workspace corev1.WorkspaceKey) {
	if dryRun {
		e.AddFieldDependencyVersionedWarning(ctx, field, name, kind, version, workspace)
	} else {
		e.AddFieldErrorf(field, MustExist, "%s %s:%s does not exist", kind, name, version)
	}
}

// AddFieldDependencyVersionedWarning adds a validation warning as an error
func (e *Error) AddFieldDependencyVersionedWarning(ctx context.Context, field string, name, kind string, version corev1.ObjectVersion, workspace corev1.WorkspaceKey) {
	warning := Warning{
		WarningType: WarningTypeDependency,
		Kind:        kind,
		Name:        name,
		Version:     version,
		Workspace:   workspace,
	}

	e.AddFieldError(field, ShouldExist, warningMessage(ctx, warning))
}

// AddFieldDeprecationWarning adds a validation warning as an error
func (e *Error) AddFieldDeprecationWarning(ctx context.Context, field string, kind, apiVersion string) {
	warning := Warning{
		WarningType: WarningTypeFieldDeprecated,
		Kind:        kind,
		Name:        field,
		APIVersion:  apiVersion,
	}

	e.AddFieldError(field, Deprecated, warningMessage(ctx, warning))
}

func warningMessage(ctx context.Context, warning Warning) string {
	warningBytes, err := json.Marshal(warning)

	var warningMessage string
	if err != nil {
		common.Log(ctx).WithError(err).Error("failed to marshal warning")
		warningMessage = "{\"warningType\":\"\",\"kind\":\"\",\"name\":\"\",\"workspace\":\"\"}"
	} else {
		warningMessage = string(warningBytes)
	}

	return warningMessage
}

// AddFieldErrorsIfPresent adds any field errors if present on any error object
// - will return true if a validation error WITH field errors
// - will return false if a validation error with no field errors present
func (e *Error) AddFieldErrorsIfPresent(err error) bool {
	v, ok := err.(*Error)
	if ok {
		if v.HasErrors() {
			e.FieldErrors = append(e.FieldErrors, v.FieldErrors...)
			return true
		}
	}
	return false
}

// AddFieldErrorf adds an error for a specific field to a validation error.
func (e *Error) AddFieldErrorf(field string, errCode ErrorCode, format string, args ...interface{}) *Error {
	e.FieldErrors = append(e.FieldErrors, NewFieldErrorf(field, errCode, format, args...))

	return e
}

// Append is called to apply aggregate any field errors from another error, ignoring duplicates
func (e *Error) Append(err *Error) {
	if err == nil || len(err.FieldErrors) == 0 {
		return
	}

	for _, fe := range err.FieldErrors {
		e.AddNewFieldError(fe)
	}
}

// ContainsFieldError checks if an error with a matching field exists
func (e *Error) ContainsFieldError(field string) bool {
	for _, e := range e.FieldErrors {
		if e.Field == field {
			return true
		}
	}
	return false
}

// FieldError provides information about a validation error on a specific field.
type FieldError struct {
	// Field causing the error, in format x.y.z
	Field string `json:"field"`
	// ErrCode is the type of constraint which has been broken.
	ErrCode ErrorCode `json:"errCode"`
	// Message is a human-readable description of the validation error.
	Message string `json:"message"`
}

// ErrorCode is the type of validation error detected.
type ErrorCode string

// The error codes should match the validator names from JSON Schema
const (
	// Deprecated indicates the supplied value is deprecated
	Deprecated ErrorCode = "deprecated"
	// MinLength error indicates the supplied value is shorter than the allowed minimum.
	MinLength ErrorCode = "minLength"
	// MaxLength error indicates the supplied value is longer than the allowed maximum.
	MaxLength ErrorCode = "maxLength"
	// Required error indicates that a field must be specified.
	Required ErrorCode = "required"
	// Pattern error indicates the input doesn't match the required regex pattern
	Pattern ErrorCode = "pattern"
	// MustExist error indicates that the named reference must exist
	MustExist ErrorCode = "mustExist"
	// ShouldExist error indicates that the named reference should exist
	ShouldExist ErrorCode = "shouldExist"
	// ReadOnly error indicates that the given value cannot be changed from a pre-defined value
	ReadOnly ErrorCode = "readOnly"
	// InvalidType error indicates that we've expected a different type
	InvalidType ErrorCode = "invalidType"
	// InvalidValue error indicates that the given value is invalid
	InvalidValue ErrorCode = "invalidValue"
	// NotAllowed error indicates that the given value is not allowed
	NotAllowed ErrorCode = "notAllowed"
	// MustBeUnique error indicates that the given value is not unique
	MustBeUnique ErrorCode = "mustBeUnique"
	// Immutable error indicates that the specified field cannot be changed after creation
	Immutable ErrorCode = "immutable"
	// FieldWarning error indicates that the given value is a warning
	FieldWarning ErrorCode = "fieldWarning"
	// NotYetImplemented error indicates that the given value is not currently
	// supported, but will be supported at some point in future.
	NotYetImplemented ErrorCode = "notYetImplemented"
)

func TranslateErrorCodeToCauseType(code ErrorCode) metav1.CauseType {
	switch code {
	case MinLength, MaxLength, Pattern, InvalidType, InvalidValue, ReadOnly, Immutable:
		return metav1.CauseTypeFieldValueInvalid
	case Required:
		return metav1.CauseTypeFieldValueRequired
	case MustExist:
		return metav1.CauseTypeFieldValueNotFound
	case NotAllowed:
		return metav1.CauseTypeFieldValueNotSupported
	case MustBeUnique:
		return metav1.CauseTypeFieldValueDuplicate
	default:
		return metav1.CauseTypeUnexpectedServerResponse
	}
}

func TranslateCauseTypeErrorCodeTo(cause metav1.CauseType) ErrorCode {
	switch cause {
	case metav1.CauseTypeFieldValueInvalid:
		return InvalidValue
	case metav1.CauseTypeFieldValueDuplicate:
		return MustBeUnique
	case metav1.CauseTypeFieldValueRequired:
		return Required
	case metav1.CauseTypeFieldValueNotFound:
		return MustExist
	case metav1.CauseTypeFieldValueNotSupported:
		return NotAllowed
	default:
		return "unknown"
	}
}

// APIStatusToValidationError converts the given error to a validation error if it's a matching API Status error
func APIStatusToValidationError(err error) error {
	if status := kerrors.APIStatus(nil); errors.As(err, &status) {
		details := status.Status().Details
		switch status.Status().Reason {
		case metav1.StatusReasonInvalid:
			ve := Error{}
			if details != nil {
				ve.Message = fmt.Sprintf("%s.%s %q is invalid", details.Kind, details.Group, details.Name)
				for _, c := range status.Status().Details.Causes {
					ve.AddFieldError(c.Field, TranslateCauseTypeErrorCodeTo(c.Type), c.Message)
				}
				return ve
			}
		case metav1.StatusReasonConflict:
			ve := ErrDependencyViolation{}
			found := false
			if details != nil {
				for _, c := range details.Causes {
					if c.Type != CauseTypeDependencyViolation {
						continue
					}
					ref, err := DependentReferenceFromString(c.Message)
					if err != nil {
						continue
					}
					ve.Dependents = append(ve.Dependents, ref)
					found = true
				}
			}
			if found {
				return ve
			}
		}
	}

	return err
}
