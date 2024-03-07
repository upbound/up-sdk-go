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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpaceMode is the space mode.
type SpaceMode string

const (
	// ModeConnected represents a space connected via connect agent.
	ModeConnected SpaceMode = "connected"
	// ModeLegacy represents a legacy space.
	ModeLegacy SpaceMode = "legacy"
	// ModeManaged represents an Upbound managed space.
	ModeManaged SpaceMode = "managed"
)

// SpaceSpec is space's spec.
type SpaceSpec struct {
	Mode SpaceMode `json:"mode"`
}

// SpaceStatus is space's status.
type SpaceStatus struct{}

// +kubebuilder:object:root=true

// A Space represents Upbound Space.
// +kubebuilder:printcolumn:name="SPACES VERSION",type="string",JSONPath=".spec.spacesConfig.version"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories=claim
type Space struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   SpaceSpec   `json:"spec"`
	Status SpaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpaceList contains a list of Spaces.
type SpaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Space `json:"items"`
}

var (
	// SpaceKind is the kind of Space.
	SpaceKind = reflect.TypeOf(Space{}).Name()
)

func init() {
	SchemeBuilder.Register(&Space{}, &SpaceList{})
}
