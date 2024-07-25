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

	"github.com/upbound/up-sdk-go/apis/common"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	spacesv1alpha1 "github.com/upbound/up-sdk-go/apis/spaces/v1alpha1"
)

// SharedTelemetryConfigAnnotation is the annotation used to mark a controlplane
// or OpenTelemetryCollector as managed by a SharedTelemetryConfig.
const SharedTelemetryConfigAnnotation = "spaces.upbound.io/sharedtelemetryconfig"

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Selected",type="string",JSONPath=`.metadata.annotations.sharedtelemetryconfig\.internal\.spaces\.upbound\.io/selected`
// +kubebuilder:printcolumn:name="Failed",type="string",JSONPath=`.metadata.annotations.sharedtelemetryconfig\.internal\.spaces\.upbound\.io/failed`
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=`.metadata.annotations.sharedtelemetryconfig\.internal\.spaces\.upbound\.io/provisioned`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories=observability,shortName=stc

// SharedTelemetryConfig defines a telemetry configuration over a set of ControlPlanes.
type SharedTelemetryConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SharedTelemetryConfigSpec   `json:"spec"`
	Status SharedTelemetryConfigStatus `json:"status,omitempty"`
}

// SharedTelemetryConfigList contains a list of SharedTelemetryConfigs.
//
// +kubebuilder:object:root=true
type SharedTelemetryConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedTelemetryConfig `json:"items"`
}

// Objects return the list of items.
func (s *SharedTelemetryConfigList) Objects() []client.Object {
	var objs = make([]client.Object, len(s.Items))
	for i := range s.Items {
		objs[i] = &s.Items[i]
	}
	return objs
}

// SharedTelemetryConfigSpec defines a telemetry configuration over a set of ControlPlanes.
type SharedTelemetryConfigSpec struct {
	// ControlPlaneSelector defines the selector for ControlPlanes on which to
	// configure telemetry.
	// +kubebuilder:validation:XValidation:rule="(has(self.labelSelectors) || has(self.names)) && (size(self.labelSelectors) > 0 || size(self.names) > 0)",message="either names or a labelSelector must be specified"
	ControlPlaneSelector spacesv1alpha1.ResourceSelector `json:"controlPlaneSelector"`

	// Exporters defines the exporters to configure on the selected ControlPlanes.
	// Untyped as we use the underlying OpenTelemetryOperator to configure the
	// OpenTelemetry collector's exporters. Use the OpenTelemetry Collector
	// documentation to configure the exporters.
	// Currently only supported exporters are push based exporters.
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Required
	Exporters common.JSONObject `json:"exporters"`

	// ExportPipeline defines the telemetry exporter pipeline to configure on
	// the selected ControlPlanes.
	// +kubebuilder:validation:Required
	ExportPipeline Pipeline `json:"exportPipeline,omitempty"`
}

// Pipeline defines the telemetry exporter pipeline to configure on the
// selected ControlPlanes.
type Pipeline struct {
	// Metrics defines the metrics exporter pipeline to configure on the
	// selected ControlPlanes. The value has to be present in the
	// spec.exporters field.
	// +kubebuilder:validation:MaxItems=10
	// +optional
	Metrics []string `json:"metrics,omitempty"`
	// Traces defines the traces exporter pipeline to configure on the
	// selected ControlPlanes. The value has to be present in the
	// spec.exporters field.
	// +kubebuilder:validation:MaxItems=10
	// +optional
	Traces []string `json:"traces,omitempty"`
	// Logs defines the logs exporter pipeline to configure on the
	// selected ControlPlanes. The value has to be present in the
	// spec.exporters field.
	// +kubebuilder:validation:MaxItems=10
	// +optional
	Logs []string `json:"logs,omitempty"`
}

// SharedTelemetryConfigStatus represents the observed state of a SharedTelemetryConfig.
type SharedTelemetryConfigStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// SelectedControlPlanes represents the names of the selected ControlPlanes.
	// +optional
	// +listType=set
	SelectedControlPlanes []string `json:"selectedControlPlanes,omitempty"`

	// list of provisioning failures.
	// +optional
	// +listType=map
	// +listMapKey=controlPlane
	Failed []SharedTelemetryConfigProvisioningFailure `json:"failed,omitempty"`

	// List of successfully provisioned targets.
	// +optional
	// +listType=set
	Provisioned []string `json:"provisioned,omitempty"`
}

// SharedTelemetryConfigProvisioningFailure defines configuration provisioning failure.
type SharedTelemetryConfigProvisioningFailure struct {
	xpv1.ResourceStatus `json:",inline"`

	// ControlPlane name where the failure occurred.
	ControlPlane string `json:"controlPlane"`
}

const (
	// ConditionTypeFailed indicates that the controlplane telemetry
	// provisioning has failed.
	ConditionTypeFailed xpv1.ConditionType = "Failed"
)

const (
	// ReasonSelectorConflict indicates that the controplane was selected by
	// multiple SharedTelemetryConfigs.
	ReasonSelectorConflict xpv1.ConditionReason = "SelectorConflict"
	// ReasonInvalidConfig indicates that the telemetry configuration is
	// invalid.
	ReasonInvalidConfig xpv1.ConditionReason = "InvalidTelemetryConfig"
)

// SelectorConflict returns a condition that indicates the controlplane is
// selected by multiple SharedTelemetryConfigs.
func SelectorConflict(msg string) xpv1.Condition {
	return xpv1.Condition{
		Type:               ConditionTypeFailed,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSelectorConflict,
		Message:            msg,
	}
}

// InvalidConfig returns a condition that indicates the controlplane
// OpenTelemetry configuration is invalid.
func InvalidConfig(msg string) xpv1.Condition {
	return xpv1.Condition{
		Type:               ConditionTypeFailed,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonInvalidConfig,
		Message:            msg,
	}
}

var (
	// SharedTelemetryConfigKind is the kind of a SharedTelemetryConfig.
	SharedTelemetryConfigKind = reflect.TypeOf(SharedTelemetryConfig{}).Name()
)

func init() {
	SchemeBuilder.Register(&SharedTelemetryConfig{}, &SharedTelemetryConfigList{})
}
