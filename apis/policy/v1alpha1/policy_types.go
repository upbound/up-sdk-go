// Copyright 2023 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	"reflect"

	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/up-sdk-go/apis/spaces/v1alpha1"
)

// SharedUpboundPolicy specifies a shared Kyverno policy projected into the specified
// ControlPlanes of the same namespace as SharedUpboundPolicy.
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Provisioned",type=string,JSONPath=`.metadata.annotations.sharedupboundpolicies\.internal\.spaces\.upbound\.io/provisioned-total`
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories={policies},shortName=sup
type SharedUpboundPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SharedUpboundPolicySpec   `json:"spec,omitempty"`
	Status SharedUpboundPolicyStatus `json:"status,omitempty"`
}

// SharedUpboundPolicyList contains a list of SharedUpboundPolicy.
// +kubebuilder:object:root=true
type SharedUpboundPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedUpboundPolicy `json:"items"`
}

// Objects return the list of items.
func (s *SharedUpboundPolicyList) Objects() []client.Object {
	var objs = make([]client.Object, len(s.Items))
	for i := range s.Items {
		objs[i] = &s.Items[i]
	}
	return objs
}

// SharedUpboundPolicySpec defines the desired state of SharedUpboundPolicy.
// +kubebuilder:validation:XValidation:rule="has(self.policyName) == has(oldSelf.policyName)",message="policyName is immutable"
type SharedUpboundPolicySpec struct {
	// PolicyName is the name to use when creating policy within a control plane.
	// optional, if not set, SharedUpboundPolicy name will be used.
	// When set, it is immutable.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="policyName is immutable"
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:MinLength=1
	// +optional
	PolicyName string `json:"policyName,omitempty"`

	// The metadata of the policy to be created.
	// +optional
	PolicyMetadata *v1alpha1.ResourceMetadata `json:"policyMetadata,omitempty"`

	// The policy is projected only to control planes
	// matching the provided selector. Either names or a labelSelector must be specified.
	// +kubebuilder:validation:XValidation:rule="(has(self.labelSelectors) || has(self.names)) && (size(self.labelSelectors) > 0 || size(self.names) > 0)",message="either names or a labelSelector must be specified"
	ControlPlaneSelector v1alpha1.ResourceSelector `json:"controlPlaneSelector"`

	// The rest of spec follows Kyverno policy spec.
	// See https://htmlpreview.github.io/?https://github.com/kyverno/kyverno/blob/main/docs/user/crd/index.html#kyverno.io/v1.Spec
	kyvernov1.Spec `json:",inline"`
}

// SharedUpboundPolicyStatus defines the observed state of the projected polcies.
type SharedUpboundPolicyStatus struct {
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
	Failed []SharedUpboundPolicyProvisioningFailure `json:"failed,omitempty"`

	// List of successfully provisioned targets.
	// +optional
	// +listType=map
	// +listMapKey=controlPlane
	Provisioned []SharedUpboundPolicyProvisioningSuccess `json:"provisioned,omitempty"`
}

// SharedUpboundPolicyProvisioningFailure defines policy provisioning failure.
type SharedUpboundPolicyProvisioningFailure struct {
	// ControlPlane name where the failure occurred.
	ControlPlane string `json:"controlPlane"`

	// List of conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// SharedUpboundPolicyProvisioningSuccess defines policy provisioning success.
type SharedUpboundPolicyProvisioningSuccess struct {
	// ControlPlane name where the external secret got successfully projected.
	ControlPlane string `json:"controlPlane"`
}

// ClusterPolicyConditionType is controlplane ClusterPolicy condition type
type ClusterPolicyConditionType string

// ClusterPolicyStatusCondition is controlplane ClusterPolicy status condition
type ClusterPolicyStatusCondition struct {
	Type   ClusterPolicyConditionType `json:"type"`
	Status corev1.ConditionStatus     `json:"status"`

	// +optional
	Message string `json:"message,omitempty"`
}

// ControlPlaneSelector returns a function that can be used for checking
// if a given object matches the selector.
func (c *SharedUpboundPolicy) ControlPlaneSelector() func(obj client.Object) (bool, error) {
	return func(obj client.Object) (bool, error) {
		return c.Spec.ControlPlaneSelector.Matches(obj)
	}
}

var (
	// SharedUpboundPolicyKind is kind of SharedUpboundPolicy
	SharedUpboundPolicyKind = reflect.TypeOf(SharedUpboundPolicy{}).Name()

	// SharedUpboundPolicyGroupKind is group kind of SharedUpboundPolicy
	SharedUpboundPolicyGroupKind = schema.GroupKind{Group: Group, Kind: SharedUpboundPolicyKind}.String()

	// SharedUpboundPolicyKindAPIVersion is apiVersion and kind of SharedUpboundPolicy
	SharedUpboundPolicyKindAPIVersion = SharedUpboundPolicyKind + "." + SchemeGroupVersion.String()

	// SharedUpboundPolicyGroupVersionKind is GVK of SharedUpboundPolicy
	SharedUpboundPolicyGroupVersionKind = SchemeGroupVersion.WithKind(SharedUpboundPolicyKind)

	// SharedUpboundPolicyGroupVersionResource is GVR of SharedUpboundPolicy
	SharedUpboundPolicyGroupVersionResource = SchemeGroupVersion.WithResource("sharedupboundpolicies")
)

func init() {
	SchemeBuilder.Register(&SharedUpboundPolicy{}, &SharedUpboundPolicyList{})
}
