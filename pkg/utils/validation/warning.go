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
	"fmt"

	corev1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
)

type WarningType string

const (
	WarningTypeDependency      WarningType = "Dependency"
	WarningTypeGeneral         WarningType = "General"
	WarningTypeFieldDeprecated WarningType = "FieldDeprecated"
)

const (
	WarningHeader = "warning"
)

type Warning struct {
	WarningType WarningType          `json:"warningType"`
	APIVersion  string               `json:"apiVersion,omitempty"`
	Kind        string               `json:"kind,omitempty"`
	Name        string               `json:"name,omitempty"`
	Version     corev1.ObjectVersion `json:"version,omitempty"`
	Workspace   corev1.WorkspaceKey  `json:"workspace,omitempty"`
	Message     string               `json:"message,omitempty"`
}

func (w Warning) shouldHandleWarning() bool {
	return (w.Name != "" && w.Kind != "") || w.Message != ""
}

func (w Warning) GetDisplayMessage() string {
	if !w.shouldHandleWarning() {
		return ""
	}

	switch w.WarningType {
	case WarningTypeDependency:
		name := w.Name
		if w.Version != "" {
			name = fmt.Sprintf("%s (version %s)", name, w.Version)
		}
		if w.Workspace == "" {
			return fmt.Sprintf("Dependency %s %s does not exist", w.Kind, name)
		} else {
			return fmt.Sprintf("Dependency %s %s %s does not exist", w.Kind, w.Workspace, name)
		}
	case WarningTypeFieldDeprecated:
		return getDeprecationWarning(w.Name, w.APIVersion, w.Kind)
	case WarningTypeGeneral:
		return fmt.Sprintf("* %s: %s", w.Name, w.Message)
	default:
		return ""
	}
}

func getDeprecationWarning(name, version, kind string) string {
	versionKind := kind
	if version != "" {
		versionKind = fmt.Sprintf("%s/%s", version, kind)
	}

	return fmt.Sprintf("Field %s on %s is deprecated and will be removed in a later version", name, versionKind)
}
