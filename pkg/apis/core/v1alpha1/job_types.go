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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type JobStatus struct {
	// RunStatus shows the history of recent reconciliation job running operations.
	RunStatus []RunStatus `json:"runStatus,omitempty"`
	// InProgressRunID is the ID of the current reconciliation operation in progress.
	InProgressRunID string `json:"inProgressRunID,omitempty"`
	// InProgressDeleteRunID is the ID of the current delete operation in progress.
	InProgressDeleteRunID string `json:"inProgressDeleteRunID,omitempty"`
	// LastProcessed is the information about the last processed attempt.
	LastProcessed LastProcessedInfo `json:"lastProcessed,omitempty"`
	// LastProcessedDelete is the information about the last processed delete attempt.
	LastProcessedDelete LastProcessedInfo `json:"lastProcessedDelete,omitempty"`
}

// AddRunState will add a new run to the status. If the run already exists, it will be reset to the
// provided details.
func (s *JobStatus) AddRunState(runID string, state RunState, isDelete bool, startTime metav1.Time) {
	rs := RunStatus{
		RunID:           runID,
		StartTime:       startTime,
		LastCheckedTime: startTime,
		IsDelete:        isDelete,
		State:           state,
	}
	// reset it if already present
	for i := range s.RunStatus {
		if s.RunStatus[i].RunID == runID {
			s.RunStatus[i] = rs
			return
		}
	}
	// else add it
	s.RunStatus = append(s.RunStatus, rs)
}

// SetRunState will set the state of an existing run to the provided details. This will no-op if the
// run does not exist.
func (s *JobStatus) SetRunState(runID string, state RunState, lastCheckedTime metav1.Time, finishTime *metav1.Time) {
	for i := range s.RunStatus {
		if s.RunStatus[i].RunID == runID {
			s.RunStatus[i].State = state
			s.RunStatus[i].FinishTime = finishTime
		}
	}
}

func (s *JobStatus) GetRunState(runID string) *RunStatus {
	for i := range s.RunStatus {
		if s.RunStatus[i].RunID == runID {
			return &s.RunStatus[i]
		}
	}
	return nil
}

type LastProcessedInfo struct {
	// Generation is the generation of the last processed attempt. If the generation is
	// different to this number (i.e. the spec has changed), the resource will be re-processed.
	Generation int64 `json:"generation,omitempty"`
	// RunID is the ID of the last run
	RunID string `json:"runID,omitempty"`
}

type RunStatus struct {
	// RunID is the unique identifier for the reconciliation run. The value is provider-specific.
	RunID string `json:"runID"`
	// IsDelete indicates whether this run is a delete reconciliation.
	// +kubebuilder:validation:Optional
	IsDelete bool `json:"isDelete,omitempty"`
	// State is the current state of the reconciliation run.
	State RunState `json:"state"`
	// StartTime is when this run started.
	StartTime metav1.Time `json:"startTime"`
	// LastCheckedTime is when the status of this run was last checked.
	// +kubebuilder:validation:Optional
	LastCheckedTime metav1.Time `json:"lastCheckedTime,omitempty"`
	// FinishTime is when this run finished, if completed/failed.
	// +kubebuilder:validation:Optional
	FinishTime *metav1.Time `json:"finishTime,omitempty"`
}

type RunState string

const (
	RunStateInProgress RunState = "InProgress"
	RunStateSucceeded  RunState = "Succeeded"
	RunStateFailed     RunState = "Failed"
)
