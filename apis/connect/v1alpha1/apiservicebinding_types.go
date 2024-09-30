// Copyright 2024 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// APIServiceBinding binds an API service represented by an APIServiceExport
// in an Upbound Space into a consumer cluster.
// This object lives in the consumer cluster.
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Versions",type="string",JSONPath=`.metadata.annotations.connect\.upbound\.io/versions`
// +kubebuilder:printcolumn:name="Export",type="string",JSONPath=`.spec.export.name`
// +kubebuilder:printcolumn:name="Namespace Binding",type="string",JSONPath=`.spec.isolation.namespaceIdentity`
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=`.status.message`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={connect},shortName={binding,bindings}
type APIServiceBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +required
	Spec   APIServiceBindingSpec   `json:"spec,omitempty"`
	Status APIServiceBindingStatus `json:"status,omitempty"`
}

// APIServiceBindingList contains a list of APIServiceBindings.
// +kubebuilder:object:root=true
type APIServiceBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIServiceBinding `json:"items"`
}

// APIServiceBindingSpec defines the configuration for an APIServiceBinding.
type APIServiceBindingSpec struct {
	// kubeconfigSecretRef is a reference to a secret containing the kubeconfig
	// to access the APIServiceExports in the Upbound Space. The user setting this
	// field needs verb=get permissions to the referenced secret.
	// Secrets should be named with the UID of the ClusterBinding
	// that was created at the Upbound Space for the particular
	// app cluster.
	// +required
	KubeconfigSecretRef xpv1.SecretReference `json:"kubeconfigSecretRef,omitempty"`
	// export is a reference of a target APIServiceExport to bind to.
	// +required
	Export APIServiceExportRef `json:"export,omitempty"`
	// request is a reference to an APIServiceBindingRequest at the service
	// provider Upbound Space. It is unset first, then set by the agent
	// afterward.
	Request APIServiceBindingRequestRef `json:"request,omitempty"`
	// clusterBinding is a reference to a ClusterBinding at the service
	// provider Upbound Space. This must match the name of the ClusterBinding
	// provisioned for the consumer cluster, at the service provider Upbound Space.
	// +required
	ClusterBinding ClusterBindingRef `json:"clusterBinding,omitempty"`
	// isolation controls how the bound API will be consumed identity-wise
	// +kubebuilder:default={}
	Isolation IsolationConfig `json:"isolation,omitempty"`
}

// APIServiceBindingRequestRef identifies an APIServiceBindingRequest in the
// service provider Upbound Space.
type APIServiceBindingRequestRef struct {
	// uid is the uid of the APIServiceBindingRequest object
	UID types.UID `json:"uid,omitempty"`
	// name is the name of the APIServiceBindingRequest object
	Name string `json:"name,omitempty"`
	// group is the Space group of the APIServiceBindingRequest object
	Group string `json:"group,omitempty"`
}

// ClusterBindingRef identifies a ClusterBinding in the
// service provider Upbound Space.
type ClusterBindingRef struct {
	// name is the name of the ClusterBinding object to reference
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name,omitempty"`
}

// IsolationConfig configures the identity restrictions for the consumption
// of the bound API
type IsolationConfig struct {
	// namespaceIdentity configures whether per-namespace
	// identities via NamespaceBinding objects are required.
	// when consuming bound APIs.
	// Valid values are: Required, Allowed, Disallowed
	// +optional
	// +kubebuilder:validation:Enum=Allowed;Disallowed;Required
	// +kubebuilder:default=Allowed
	NamespaceIdentity NamespaceIdentityMode `json:"namespaceIdentity,omitempty"`
}

// NamespaceIdentityMode describes the NamespaceBinding requirement mode.
type NamespaceIdentityMode string

const (
	// NamespaceIdentityAllowed allows per-namespace identities in consumer clusters
	// but does not enforce
	NamespaceIdentityAllowed NamespaceIdentityMode = "Allowed"
	// NamespaceIdentityDisallowed prohibits the usage of per-namespace identities.
	// all claims of a particular bound API in the consumer cluster is consumed
	// via the common agent credentials of the consumer cluster.
	NamespaceIdentityDisallowed NamespaceIdentityMode = "Disallowed"
	// NamespaceIdentityRequired enforces per-namespace identity to be present,
	// via the existence of NamespaceBinding
	NamespaceIdentityRequired NamespaceIdentityMode = "Required"
)

// APIServiceBindingStatus defines the status of a APIServiceBinding.
// Reports the APIServiceExport that was bound to, and the success status.
type APIServiceBindingStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// Message provides human-readable information about the current status of
	// the APIServiceBinding.
	// +optional
	Message string `json:"message,omitempty"`
}

var (
	// APIServiceBindingKind is kind of APIServiceBinding
	APIServiceBindingKind = reflect.TypeOf(APIServiceBinding{}).Name()
)

func init() {
	SchemeBuilder.Register(&APIServiceBinding{}, &APIServiceBindingList{})
}
