// Copyright 2024 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

// ClusterBinding represents a bound consumer cluster.
// It lives in the Upbound Spaces that provides the APIServiceExports.
//
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Description",type="string",JSONPath=`.status.description`
// +kubebuilder:printcolumn:name="Agent",type="string",JSONPath=`.status.agentVersion`
// +kubebuilder:printcolumn:name="Heartbeat",type="date",JSONPath=`.status.lastHeartbeatTime`
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=`.status.message`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={connect}
type ClusterBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +required
	Spec   ClusterBindingSpec   `json:"spec,omitempty"`
	Status ClusterBindingStatus `json:"status,omitempty"`
}

// ClusterBindingList contains a list of ClusterBindings.
// +kubebuilder:object:root=true
type ClusterBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterBinding `json:"items"`
}

// ClusterBindingSpec defines the desired state of a ClusterBinding.
type ClusterBindingSpec struct{}

// ClusterBindingStatus stores status information about a service binding. It is
// updated by both the konnector and the service provider.
type ClusterBindingStatus struct {
	// conditions is a list of conditions that apply to the ClusterBinding. It is
	// updated by the konnector agent and the service provider.
	xpv1.ResourceStatus `json:",inline"`
	// description is the human-readable description of the consumer cluster.
	// +optional
	Description string `json:"description,omitempty"`
	// lastHeartbeatTime is the last time the konnector updated the status.
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// heartbeatInterval is the maximal interval between heartbeats that the
	// konnector promises to send. The service provider can assume that the
	// konnector is not unhealthy if it does not receive a heartbeat within
	// this time.
	HeartbeatInterval metav1.Duration `json:"heartbeatInterval,omitempty"`
	// konnectorVersion is the version of the konnector that is running on the
	// consumer cluster.
	AgentVersion string `json:"agentVersion,omitempty"`

	// Message provides human-readable information about the current status of
	// the ClusterBinding.
	// +optional
	Message string `json:"message,omitempty"`
}

var (
	// ClusterBindingKind is kind of ClusterBinding
	ClusterBindingKind = reflect.TypeOf(ClusterBinding{}).Name()
)

func init() {
	SchemeBuilder.Register(&ClusterBinding{}, &ClusterBindingList{})
}
