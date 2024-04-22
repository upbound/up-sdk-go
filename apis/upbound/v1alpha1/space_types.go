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

const (
	// SpaceModeLabelKey is the key used to identify the connection mode to
	// Upbound. The value should be of type `SpaceMode`.
	SpaceModeLabelKey = "spaces.upbound.io/mode"

	// SpaceRegionLabelKey is the key used to identify the cloud region for the
	// space. The region comes from the Upbound list of regions, independent
	// from any cloud provider region list. The value should always match the
	// `spec.region` field.
	SpaceRegionLabelKey = "spaces.upbound.io/region"

	// SpaceProviderLabelKey is the key used to identify the cloud provider for
	// the space. The value should always match the `spec.provider` field.
	SpaceProviderLabelKey = "spaces.upbound.io/provider"
)

// SpaceMode is the mode in which the space connects to Upbound.
type SpaceMode string

const (
	// ModeConnected represents a space connected via connect agent.
	ModeConnected SpaceMode = "connected"
	// ModeLegacy represents a legacy space.
	ModeLegacy SpaceMode = "legacy"
	// ModeManaged represents an Upbound managed space.
	ModeManaged SpaceMode = "managed"
)

// CloudProvider is the hosting cloud provider for the space.
type CloudProvider string

const (
	// CloudProviderGCP represents the space lives on GCP.
	CloudProviderGCP CloudProvider = "gcp"
	// CloudProviderAWS represents the space lives on AWS.
	CloudProviderAWS CloudProvider = "aws"
	// CloudProviderUnknown represents the space lives in an unknown provider.
	CloudProviderUnknown CloudProvider = "unknown"
)

// Region is the region in which the space is hosted.
type Region string

const (
	// RegionUSWest1 represents the space lives in US-West-1 of its respective provider.
	RegionUSWest1 Region = "us-west-1"
	// RegionUSEast1 represents the space lives in US-East-1 of its respective provider.
	RegionUSEast1 Region = "us-east-1"
	// RegionUSCentral1 represents the space lives in US-Central-1 of its respective provider.
	RegionUSCentral1 Region = "us-central-1"
)

// ConnectionStatus represents the ability for clients to be able to connect to
// the space.
type ConnectionStatus string

const (
	// ConnectionStatusConnected represents the space is reachable.
	ConnectionStatusConnected ConnectionStatus = "connected"
	// ConnectionStatusUnreachable represents the space cannot currently be
	// reached.
	ConnectionStatusUnreachable ConnectionStatus = "unreachable"
	// ConnectionStatusUnknown represents the space is not known.
	ConnectionStatusUnknown ConnectionStatus = "unknown"
)

// ConnectionDetails is the collection of statuses and timestamps surrounding
// the connection to the space.
type ConnectionDetails struct {
	// +kubebuilder:default="unknown"
	// The current status of the connection to the space
	Status ConnectionStatus `json:"status,omitempty"`
}

// SpaceSpec is space's spec.
type SpaceSpec struct {
	Provider *CloudProvider `json:"provider,omitempty"`
	Region   *Region        `json:"region,omitempty"`
}

// SpaceStatus is space's status.
type SpaceStatus struct {
	// +optional
	// The FQDN endpoint for the Space Cluster used for Ingress
	FQDN string `json:"fqdn,omitempty"`
	// The statuses and timestamps surrounding the connection to the space
	ConnectionDetails ConnectionDetails `json:"connection,omitempty"`
}

// +kubebuilder:object:root=true

// A Space is a kubernetes style representation of an Upbound Space.
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.connection.status"
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
