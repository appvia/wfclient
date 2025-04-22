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

type ScopeWorkspace struct {
	// AllWorkspaces indicates all workspaces can use this resource
	// +kubebuilder:validation:Optional
	AllWorkspaces bool `json:"allWorkspaces,omitempty"`
	// AllowedWorkspaces indicates which workspaces can use this resource.
	// only valid if All is false
	// +kubebuilder:validation:Optional
	AllowedWorkspaces WorkspaceKeys `json:"allowedWorkspaces,omitempty" type:"[]string"`
}

func (s ScopeWorkspace) IsSet() bool {
	return s.AllWorkspaces || len(s.AllowedWorkspaces) > 0
}

func (s ScopeWorkspace) ToAllocationPtr() *ResourceAllocation {
	v := s.ToAllocation()
	return &v
}

func (s ScopeWorkspace) ToAllocation() ResourceAllocation {
	r := ResourceAllocation{
		Type: ResourceAllocationNone,
	}
	if s.AllWorkspaces {
		r.Type = ResourceAllocationAll
		r.Workspaces = []WorkspaceKey{
			"*",
		}
	}
	if len(s.AllowedWorkspaces) > 0 {
		if !s.AllWorkspaces {
			r.Type = ResourceAllocationWorkspaces
		}
		r.Workspaces = s.AllowedWorkspaces
	}
	return r
}

func (s ScopeWorkspace) AllowsWorkspace(workspace WorkspaceKey) bool {
	if s.AllWorkspaces {
		return true
	}
	for _, ws := range s.AllowedWorkspaces {
		if ws == workspace {
			return true
		}
	}
	return false
}

type ScopeStages struct {
	// AllStages indicates all stages can use this resource
	// +kubebuilder:validation:Optional
	AllStages bool `json:"allStages,omitempty"`
	// AllowedStages indicates which stages can use this resource.
	// only valid if AllStages is false
	// +kubebuilder:validation:Optional
	AllowedStages []string `json:"allowedStages,omitempty"`
}

func (s ScopeStages) IsSet() bool {
	return s.AllStages || len(s.AllowedStages) > 0
}

func (s ScopeStages) AllowsStage(stage string) bool {
	if s.AllStages {
		return true
	}
	for _, st := range s.AllowedStages {
		if st == stage {
			return true
		}
	}
	return false
}

type Scope struct {
	// ScopeWorkspace indicates which workspaces can use this resource.
	ScopeWorkspace `json:",inline"`
	// ScopeStages indicates which stages can use this resource.
	ScopeStages `json:",inline"`
}

func (s Scope) IsSet() bool {
	return s.ScopeWorkspace.IsSet() || s.ScopeStages.IsSet()
}
