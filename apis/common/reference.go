// Copyright 2024 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

// TypedLocalObjectReference contains enough information to let you locate the
// typed referenced object inside the same namespace.
//
// +structType=atomic
// +kubebuilder:object:root=false
// +kubebuilder:object:generate=true
type TypedLocalObjectReference struct {
	// APIGroup is the group for the resource being referenced.
	// If APIGroup is not specified, the specified Kind must be in the core API group.
	// For any other third-party types, APIGroup is required.
	// +optional
	APIGroup *string `json:"apiGroup,omitempty"`

	// Kind is the type of resource being referenced
	// +kubebuilder:validation:MinLength=1
	Kind string `json:"kind,omitempty"`

	// Name is the name of resource being referenced
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name,omitempty"`
}
