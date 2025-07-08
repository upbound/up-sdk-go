// Copyright 2025 Upbound Inc
// All rights reserved

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +kubebuilder:object:root=true
// +genclient
// +genclient:nonNamespaced

// An AddOnRuntimeConfig provides settings for configuring runtime of an
// addon package.
//
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={upbound}
type AddOnRuntimeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AddOnRuntimeConfigSpec `json:"spec"`
}

// HelmConfigSpec defines the Helm-specific configuration for an addon runtime.
type HelmConfigSpec struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	Values runtime.RawExtension `json:"values,omitempty"`
}

// AddOnRuntimeConfigSpec specifies the configuration of an AddOnRuntimeConfig.
type AddOnRuntimeConfigSpec struct {
	Helm *HelmConfigSpec `json:"helm,omitempty"`
}

// +kubebuilder:object:root=true

// AddOnRuntimeConfigList contains a list of AddOnRuntimeConfig.
type AddOnRuntimeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddOnRuntimeConfig `json:"items"`
}
