// Copyright 2024 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

// NamespaceBinding defines the specific credentials to be used
// when consuming an API bound via APIServiceBinding, for claims
// to be created in a consumer cluster namespace.
// Name of this object must match to the name of APIServiceBinding
// for the desired API.
// This object lives in the consumer cluster.
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Identity Secret",type="string",JSONPath=`.spec.kubeconfigSecretRef.name`
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=`.status.message`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories={connect}
type NamespaceBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +required
	Spec   NamespaceBindingSpec   `json:"spec,omitempty"`
	Status NamespaceBindingStatus `json:"status,omitempty"`
}

// NamespaceBindingList contains a list of NamespaceBinding.
// +kubebuilder:object:root=true
type NamespaceBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceBinding `json:"items"`
}

// NamespaceBindingSpec defines the configuration for NamespaceBinding.
type NamespaceBindingSpec struct {
	// kubeconfigSecretRef is a reference to a secret that contains a
	// kubeconfig. The user setting this field needs verb=get permissions
	// to the referenced secret.
	KubeconfigSecretRef xpv1.SecretReference `json:"kubeconfigSecretRef,omitempty"`
}

// NamespaceBindingStatus defines the observed state of a NamespaceBinding.
type NamespaceBindingStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// Message provides human-readable information about the current status of
	// the NamespaceBinding.
	// +optional
	Message string `json:"message,omitempty"`
}

var (
	// NamespaceBindingKind is kind of NamespaceBinding
	NamespaceBindingKind = reflect.TypeOf(NamespaceBinding{}).Name()
)

func init() {
	SchemeBuilder.Register(&NamespaceBinding{}, &NamespaceBindingList{})
}
