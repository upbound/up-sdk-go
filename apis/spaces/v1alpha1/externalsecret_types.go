// Copyright 2023 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	"reflect"

	esv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SharedExternalSecret specifies a shared ExternalSecret projected into the specified
// ControlPlanes of the same namespace as ClusterExternalSecret and with that
// propagated into the specified namespaces.
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Provisioned",type=string,JSONPath=`.metadata.annotations.sharedexternalsecrets\.internal\.spaces\.upbound\.io/provisioned-total`
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories={externalsecrets},shortName=ses
type SharedExternalSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SharedExternalSecretSpec   `json:"spec,omitempty"`
	Status SharedExternalSecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// SharedExternalSecretList contains a list of SharedExternalSecret.
type SharedExternalSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedExternalSecret `json:"items"`
}

// Objects return the list of items.
func (s *SharedExternalSecretList) Objects() []client.Object {
	var objs = make([]client.Object, len(s.Items))
	for i := range s.Items {
		objs[i] = &s.Items[i]
	}
	return objs
}

// SharedExternalSecretSpec defines the desired state of SharedExternalSecret.
// +kubebuilder:validation:XValidation:rule="has(self.externalSecretName) == has(oldSelf.externalSecretName)",message="externalSecretName is immutable"
type SharedExternalSecretSpec struct {
	// ExternalSecretName is the name to use when creating external secret within a control plane.
	// optional, if not set, SharedExternalSecret name will be used.
	// When set, it is immutable.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="externalSecretName is immutable"
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:MinLength=1
	// +optional
	ExternalSecretName string `json:"externalSecretName,omitempty"`

	// The metadata of the secret store to be created.
	// +optional
	ExternalSecretMetadata *ResourceMetadata `json:"externalSecretMetadata,omitempty"`

	// The secret is projected only to control planes
	// matching the provided selector. Either names or a labelSelector must be specified.
	// +kubebuilder:validation:XValidation:rule="(has(self.labelSelectors) || has(self.names)) && (size(self.labelSelectors) > 0 || size(self.names) > 0)",message="either names or a labelSelector must be specified"
	ControlPlaneSelector ResourceSelector `json:"controlPlaneSelector"`

	// The projected secret can be consumed
	// only within namespaces matching the provided selector.
	// Either names or a labelSelector must be specified.
	// +kubebuilder:validation:XValidation:rule="(has(self.labelSelectors) || has(self.names)) && (size(self.labelSelectors) > 0 || size(self.names) > 0)",message="either names or a labelSelector must be specified"
	NamespaceSelector ResourceSelector `json:"namespaceSelector"`

	// The spec for the ExternalSecrets to be created.
	ExternalSecretSpec esv1beta1.ExternalSecretSpec `json:"externalSecretSpec"`

	// Used to configure secret refresh interval in seconds.
	// +optional
	RefreshInterval *metav1.Duration `json:"refreshTime,omitempty"`
}

// SharedExternalSecretStatus defines the observed state of the ExternalSecret.
type SharedExternalSecretStatus struct {
	// We needed to introduce a common field to workaround
	// https://github.com/kubernetes/kubernetes/issues/117447
	// otherwise the initial idea was that each controller
	// just updates/remove its item in the bellow lists.

	// observed resource generation.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// list of provisioning failures.
	// +optional
	// +listType=map
	// +listMapKey=controlPlane
	Failed []SharedExternalSecretProvisioningFailure `json:"failed,omitempty"`

	// List of successfully provisioned targets.
	// +optional
	// +listType=map
	// +listMapKey=controlPlane
	Provisioned []SharedExternalSecretProvisioningSuccess `json:"provisioned,omitempty"`
}

type SharedExternalSecretProvisioningFailure struct {
	// ControlPlane name where the failure occurred.
	ControlPlane string `json:"controlPlane"`

	// List of conditions.
	// +optional
	Conditions []esv1beta1.ClusterExternalSecretStatusCondition `json:"conditions,omitempty"`
}

// SharedExternalSecretProvisioningSuccess defines external secret provisioning success.
type SharedExternalSecretProvisioningSuccess struct {
	// ControlPlane name where the external secret got successfully projected.
	ControlPlane string `json:"controlPlane"`
}

// ControlPlaneSelector returns a function that can be used for checking
// if a given object matches the selector.
func (c *SharedExternalSecret) ControlPlaneSelector() func(obj client.Object) (bool, error) {
	return func(obj client.Object) (bool, error) {
		return c.Spec.ControlPlaneSelector.Matches(obj)
	}
}

var (
	SharedPlaneExternalSecretKind        = reflect.TypeOf(SharedExternalSecret{}).Name()
	SharedExternalSecretGroupKind        = schema.GroupKind{Group: Group, Kind: SharedPlaneExternalSecretKind}.String()
	SharedExternalSecretKindAPIVersion   = SharedPlaneExternalSecretKind + "." + SchemeGroupVersion.String()
	SharedExternalSecretGroupVersionKind = SchemeGroupVersion.WithKind(SharedPlaneExternalSecretKind)
)

func init() {
	SchemeBuilder.Register(&SharedExternalSecret{}, &SharedExternalSecretList{})
}
