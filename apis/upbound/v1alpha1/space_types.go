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
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type SpaceMode string

const (
	ModeConnected SpaceMode = "connected"
	ModeLegacy    SpaceMode = "legacy"
	ModeManaged   SpaceMode = "managed"
)

type CloudProvider string

const (
	CloudProviderGCP     CloudProvider = "gcp"
	CloudProviderAWS     CloudProvider = "aws"
	CloudProviderUnknown CloudProvider = "unknown"
)

type Region string

const (
	RegionUSWest1    Region = "us-west-1"
	RegionUSEast1    Region = "us-east-1"
	RegionUSCentral1 Region = "us-central-1"
)

type SpaceSpec struct {
	Mode     SpaceMode      `json:"mode"`
	Provider *CloudProvider `json:"provider,omitempty"`
	Region   *Region        `json:"region,omitempty"`
}

type SpaceStatus struct{}

// +kubebuilder:object:root=true

// A Space is a kubernetes style representation of an Upbound Space.
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

// SpaceList contains a list of Space
type SpaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Space `json:"items"`
}

// Space type metadata.
var (
	SpaceKind             = reflect.TypeOf(Space{}).Name()
	SpaceGroupKind        = schema.GroupKind{Group: Group, Kind: SpaceKind}.String()
	SpaceKindAPIVersion   = SpaceKind + "." + SchemeGroupVersion.String()
	SpaceGroupVersionKind = SchemeGroupVersion.WithKind(SpaceKind)
)

func init() {
	SchemeBuilder.Register(&Space{}, &SpaceList{})
}
