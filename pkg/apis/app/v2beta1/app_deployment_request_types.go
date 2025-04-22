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

package v2beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
)

type AppManifestFormat string

const (
	AppManifestFormatJSON AppManifestFormat = "json"
	AppManifestFormatYAML AppManifestFormat = "yaml"
)

type AppDeploymentRequest struct {
	// AppManifest is the provided JSON/YAML manifest from the user
	AppManifest string `json:"appManifest"`
	// AppManifestFormat is the format of the provided manifest ('json' or 'yaml'; default is 'yaml')
	AppManifestFormat AppManifestFormat `json:"appManifestFormat,omitempty"`
	// AppDefinitionRef is the reference to the AppDefinition to use for the deployment instead of the
	// AppManifest.
	AppDefinitionRef *corev1.PlanRef `json:"appDefinitionRef,omitempty"`
	// Vars is the set of deploy-time variables provided with the deployment request
	Vars map[string]string
	// Secrets is the set of deploy-time sensitive variables provided with the deployment request
	Secrets map[string]string
	// TargetEnv is the name of the target appenv
	TargetEnv string `json:"targetEnv"`
	// CreateAppEnv provides options to create the referenced application environment. The name of
	// the desired application environment must be specified in the TargetEnv. If the environment
	// already exists, populating this field will cause a validation error.
	CreateAppEnv *CreateAppEnvOptions `json:"createAppEnv,omitempty"`
	// CustomTimeout is the optional timeout for the deployment job. If this is not provided, a
	// system default timeout will be used. After this period, the deployment will fail if it has
	// not successfully completed.
	CustomTimeout metav1.Duration `json:"customTimeout,omitempty"`
	// Only limits the deployment/removal to the specified components
	Only []string `json:"only,omitempty"`
	// AsMe indicates if we should run as the current user instead of using deployment roles
	AsMe bool `json:"asMe,omitempty"`
	// Files is a list of files to be used to perform the deployment. For each component that has
	// a helm.chartPath populated, a HelmChart file should be added to this array. For each entry
	// in helm.additionalValuesFiles populated in each component, a HelmValuesFile should be added to
	// this array.
	Files []AppDeploymentFile `json:"files,omitempty"`
}

type AppDeploymentFileType string

const (
	AppDeploymentFileTypeHelmChart      AppDeploymentFileType = "HelmChart"
	AppDeploymentFileTypeHelmValuesFile AppDeploymentFileType = "HelmValuesFile"
)

type AppDeploymentFile struct {
	// Type is the type of file to be added.
	Type AppDeploymentFileType `json:"type"`
	// Path is the path to the file as used in the component definition.
	Path string `json:"path"`
	// Data is the data of the file. For Helm chart files, this should be a base64 encoded tar.gz.
	// For Helm values files, this should be the raw YAML data.
	Data string `json:"data"`
}

type AppDeploymentFiles []AppDeploymentFile

// Get returns the named file of the given type, or nil if it not found.
func (f AppDeploymentFiles) Get(path string, fileType AppDeploymentFileType) *AppDeploymentFile {
	for _, file := range f {
		if file.Path == path && file.Type == fileType {
			return &file
		}
	}
	return nil
}

// GetData returns the data of the named file of the given type, or an empty string if it not found.
func (f AppDeploymentFiles) GetData(path string, fileType AppDeploymentFileType) string {
	// shortcut - if no path, no point looking for it, just return empty.
	if path == "" {
		return ""
	}
	file := f.Get(path, fileType)
	if file == nil {
		return ""
	}
	return file.Data
}
