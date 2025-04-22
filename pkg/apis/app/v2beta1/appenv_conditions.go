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
	corev1alpha1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
)

const (
	// ConditionClusterReady will be set to indicate that the referenced cluster for this
	// application environment is ready for the environment to be provisioned on it.
	ConditionClusterReady corev1alpha1.ConditionType = "ClusterReady"
	// ConditionNamespaceReady will be set to indicate that the namespace for this environment is
	// provisioned and healthy on the specified cluster
	ConditionNamespaceReady corev1alpha1.ConditionType = "NamespaceReady"
	// ConditionConfigMapReady will be set to indicate that the config for this environment is
	// in the specified environment
	ConditionConfigMapReady corev1alpha1.ConditionType = "ConfigReady"
)

var AppEnvConditions = []corev1alpha1.ConditionSpec{
	{Type: ConditionClusterReady, Name: "Host cluster ready"},
	{Type: ConditionNamespaceReady, Name: "Namespace ready"},
	{Type: ConditionConfigMapReady, Name: "Config ready"},
	{Type: corev1alpha1.ConditionReady, Name: "App environment ready"},
}
