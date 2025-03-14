// Copyright 2025 Upbound Inc
// All rights reserved

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +kubebuilder:object:root=true
// +genclient
// +genclient:nonNamespaced

// An ControllerRuntimeConfig provides settings for configuring runtime of a
// controller package.
//
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={upbound}
type ControllerRuntimeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ControllerRuntimeConfigSpec `json:"spec"`
}

// HelmConfigSpec defines the Helm-specific configuration for a controller runtime.
type HelmConfigSpec struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	Values runtime.RawExtension `json:"values,omitempty"`
}

// ControllerRuntimeConfigSpec specifies the configuration of an ControllerRuntimeConfig.
type ControllerRuntimeConfigSpec struct {
	Helm *HelmConfigSpec `json:"helm,omitempty"`
}

// +kubebuilder:object:root=true

// ControllerRuntimeConfigList contains a list of ControllerRuntimeConfig.
type ControllerRuntimeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControllerRuntimeConfig `json:"items"`
}
