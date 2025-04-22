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
	corev1 "github.com/appvia/wfclient/pkg/apis/core/v1alpha1"
)

const (
	// ConditionDeploymentJobPlanned will be set when all referenced plans and environments
	// have been loaded and are ready for deployment and the plan for the job has been successfully
	// created
	ConditionDeploymentJobPlanned corev1.ConditionType = "JobPlanned"

	// ConditionDeploymentJobActionsTaken will be set when the deployment job has taken all the
	// actions on the plan.
	ConditionDeploymentJobActionsTaken corev1.ConditionType = "ActionsTaken"
)

var DeploymentJobConditions = []corev1.ConditionSpec{
	{Type: ConditionDeploymentJobPlanned, Name: "Job planned"},
	{Type: ConditionDeploymentJobActionsTaken, Name: "Deployment actions taken"},
	{Type: corev1.ConditionReady, Name: "Deployment complete"},
}
