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
	"errors"
	"fmt"
	"regexp"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
)

const CauseTypeDependencyViolation metav1.CauseType = "DependencyViolation"

type ErrDependencyViolation struct {
	Message    string               `json:"message"`
	Dependents []DependentReference `json:"dependents"`
}

func IsDependencyViolationError(err error) bool {
	if _, ok := err.(ErrDependencyViolation); ok {
		return true
	}
	return errors.As(err, &ErrDependencyViolation{})
}

func (e ErrDependencyViolation) Error() string {
	var str string
	for _, d := range e.Dependents {
		if !d.System {
			str = fmt.Sprintf("%s * %s\n", str, d.String())
		}
	}

	if str != "" {
		if e.Message == "" {
			e.Message = "the following objects need to be deleted first"
		}
		return strings.TrimRight(e.Message, ":") + ":\n" + str
	}

	for _, d := range e.Dependents {
		str = str + fmt.Sprintf(" * %s\n", d.String())
	}

	return "waiting for the following objects to be deleted by Wayfinder:\n" + str
}

// DependentReference is an object reference to a dependent object in the same namespace
type DependentReference struct {
	// Kind of the dependent
	Kind string `json:"kind"`
	// Name of the dependent
	Name string `json:"name"`
	// Version is the version of the resource, if it is a versioned resource
	Version corev1.ObjectVersion `json:"version,omitempty"`
	// Workspace of the dependant
	Workspace corev1.WorkspaceKey `json:"workspace"`
	// System is true if this is a system resource
	System bool `json:"system"`
}

func (d DependentReference) String() string {
	if d.Workspace == "" {
		if d.Version != "" {
			return fmt.Sprintf("%s/%s@%s", d.Kind, d.Name, d.Version)
		}
		return fmt.Sprintf("%s/%s", d.Kind, d.Name)
	}

	if d.Version != "" {
		return fmt.Sprintf("%s/%s/%s@%s", d.Kind, d.Workspace, d.Name, d.Version)
	}
	return fmt.Sprintf("%s/%s/%s", d.Kind, d.Workspace, d.Name)
}

// DependentReferenceFromStrong converts the string representation of a DependentReference
// back into a DependentReference
func DependentReferenceFromString(input string) (DependentReference, error) {
	d := DependentReference{}

	slashCount := strings.Count(input, "/")

	var err error
	var ws string
	if slashCount == 2 {
		err = stringScanRegex(input, "^(.+)/(.+)/(.+)$", &d.Kind, &ws, &d.Name)
		d.Workspace = corev1.WorkspaceKey(ws)
	} else if slashCount == 1 {
		err = stringScanRegex(input, "^(.+)/(.+)$", &d.Kind, &d.Name)
	} else {
		err = fmt.Errorf("incorrect dependent reference format: %s", input)
	}

	return d, err
}

func stringScanRegex(input string, expression string, output ...*string) error {
	re, err := regexp.Compile(expression)
	if err != nil {
		return err
	}
	result := re.FindStringSubmatch(input)
	if result == nil {
		return fmt.Errorf("no matches found")
	}
	if len(result) != len(output)+1 {
		return fmt.Errorf("%d output variables provided for %d submatches", len(output), len(result)-1)
	}
	for i, s := range output {
		*s = (result[i+1])
	}
	return nil
}
