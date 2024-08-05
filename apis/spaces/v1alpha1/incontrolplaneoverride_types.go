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
	"strings"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:resource:scope=Namespaced,categories=spaces

// InControlPlaneOverride represents resource configuration overrides in
// a ControlPlane. The specified override can be applied on single objects
// as well as claim/XR object hierarchies.
type InControlPlaneOverride struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InControlPlaneOverrideSpec   `json:"spec"`
	Status InControlPlaneOverrideStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InControlPlaneOverrideList is a list of InControlPlaneOverride objects.
type InControlPlaneOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InControlPlaneOverride `json:"items"`
}

// PatchPropagationPolicy denotes the traversal direction on
// an object's hierarchy.
type PatchPropagationPolicy string

const (
	// PatchPropagateAscending denotes traversal on a target object hierarchy
	// following the metadata.ownerReferences.
	PatchPropagateAscending PatchPropagationPolicy = "Ascending"
	// PatchPropagateDescending denotes traversal on a target object hierarchy
	// following the spec.resourceRef & spec.resourceRefs reference fields.
	PatchPropagateDescending PatchPropagationPolicy = "Descending"
	// PatchPropagateNone denotes that no traversal will be done and
	// only the target object will be visited.
	PatchPropagateNone PatchPropagationPolicy = "None"
)

// PatchDeletionPolicy controls what happens when an InControlPlaneOverride
// object is deleted. We either attempt to roll back the changes on the
// target object hierarchy, or we keep them.
type PatchDeletionPolicy string

const (
	// PatchDeletionRollBack attempts to roll back the changes on
	// the target object hierarchy when the InControlPlaneOverride object is
	// being deleted.
	PatchDeletionRollBack PatchDeletionPolicy = "RollBack"
	// PatchDeletionKeep keeps the changes on the target object
	// hierarchy when the InControlPlaneOverride object is being deleted.
	PatchDeletionKeep PatchDeletionPolicy = "Keep"
)

// MetadataPatch represents the Kube object metadata.
type MetadataPatch struct {
	// Annotations represents the Kube object annotations.
	// Only the following annotations are allowed to be patched:
	// - crossplane.io/paused
	// - spaces.upbound.io/force-reconcile-at
	// +kubebuilder:validation:XValidation:rule="self.all(k, k == 'crossplane.io/paused' || k == 'spaces.upbound.io/force-reconcile-at')",message="Only the crossplane.io/paused and spaces.upbound.io/force-reconcile-at annotations are allowed"
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Override represents a configuration patch which is serialized into JSON to
// obtain the fully specified intent to be used with server side apply.
type Override struct {
	// Metadata specifies the patch metadata.
	// +optional
	Metadata *MetadataPatch `json:"metadata,omitempty"`
}

// ObjectReference represents a optionally namespaces Kubernetes API object
// reference.
type ObjectReference struct {
	// APIVersion of the referenced object.
	// +kubebuilder:validation:MinLength=1
	APIVersion string `json:"apiVersion"`

	// Kind of the referenced object.
	// +kubebuilder:validation:MinLength=1
	Kind string `json:"kind"`

	// Name of the referenced object.
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// Namespace of the referenced object.
	// +optional
	Namespace *string `json:"namespace,omitempty"`
}

func (r *ObjectReference) String() string {
	if r == nil {
		return "nil"
	}
	return strings.Join([]string{
		"APIVersion: ", r.APIVersion, ", ",
		"Kind: ", r.Kind, ", ",
		"Name: ", r.Name, ", ",
		"Namespace: '", ptr.Deref(r.Namespace, ""), "'",
	}, "")
}

// InControlPlaneOverrideSpec defines a configuration override
// on a target object hierarchy in a target ControlPlane with the
// given name.
type InControlPlaneOverrideSpec struct {
	// ControlPlaneName is the name of the target ControlPlane where
	// the resource configuration overrides will be applied.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="controlPlaneName is immutable"
	ControlPlaneName string `json:"controlPlaneName"`

	// TargetRef is the object reference to a Kubernetes API object where the
	// configuration override will start. The controller will traverse the
	// target object's hierarchy depending on the PropagationPolicy. If
	// PropagationPolicy is None, then only the target object will be updated.
	TargetRef ObjectReference `json:"targetRef"`

	// PropagationPolicy specifies whether the configuration override will be
	// applied only to the object referenced in TargetRef (None), after an
	// ascending or descending hierarchy traversal will be done starting with
	// the target object.
	// +kubebuilder:validation:Enum=None;Ascending;Descending
	// +kubebuilder:default=None
	// +optional
	PropagationPolicy PatchPropagationPolicy `json:"propagationPolicy"`

	// DeletionPolicy specifies whether when the InControlPlaneOverride object
	// is deleted, the configuration override should be kept (Keep) or
	// rolled back (RollBack).
	// +kubebuilder:validation:Enum=RollBack;Keep
	// +kubebuilder:default=RollBack
	// +optional
	DeletionPolicy PatchDeletionPolicy `json:"deletionPolicy"`

	// Override denotes the configuration override to be applied on the target
	// object hierarchy. The fully specified intent is obtained by serializing
	// the Override.
	Override Override `json:"override"`
}

// PatchState denotes the result of the patch operation on the associated
// target object.
type PatchState string

const (
	// PatchStateSkipped denotes that the target object was skipped.
	// The reason for the skip is specified in the `reason` field.
	PatchStateSkipped PatchState = "Skipped"
	// PatchStateError denotes that there was a transient error while patching
	// the object.
	PatchStateError PatchState = "Error"
)

// PatchStateReason denotes why a patch operation on the associated
// target object has been skipped.
type PatchStateReason string

const (
	// PatchStateReasonConflict denotes that the patch operation on
	// the associated target object has been skipped due to a conflict with
	// another field manager.
	PatchStateReasonConflict PatchStateReason = "Conflict"
	// PatchStateReasonSchemaMismatch denotes that the patch operation on
	// the associated target object has been skipped due to a schema mismatch
	// between the fully specified intent and the object's schema.
	PatchStateReasonSchemaMismatch PatchStateReason = "SchemaMismatch"
)

// PatchedObjectStatus represents the state of an applied patch to an object
// in the target hierarchy.
type PatchedObjectStatus struct {
	// ObjectReference is the Kubernetes object reference to the object
	// which has been updated.
	ObjectReference `json:",inline"`

	// Metadata UID of the patch target object.
	// +optional
	UID *types.UID `json:"uid,omitempty"`

	// Status of the configuration override.
	// +kubebuilder:validation:Enum=Success;Skipped;Error
	Status PatchState `json:"status"`

	// Reason is the reason for the target objects override Status.
	Reason PatchStateReason `json:"reason"`

	// Message holds an optional detail message detailing the observed state.
	// +optional
	Message *string `json:"message,omitempty"`
}

// String returns a string representation of the PatchedObjectStatus.
func (r *PatchedObjectStatus) String() string {
	if r == nil {
		return "nil"
	}
	return strings.Join([]string{r.ObjectReference.String(), ", ",
		"UID: '", string(ptr.Deref(r.UID, "")), "', ",
		"Status: ", string(r.Status), ", ",
		"Reason: ", string(r.Reason), ", ",
		"Message: '", ptr.Deref(r.Message, ""), "'"}, "")
}

// InControlPlaneOverrideStatus defines the status of an InControlPlaneOverride
// object.
type InControlPlaneOverrideStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// +optional
	ObjectRefs []PatchedObjectStatus `json:"objectRefs,omitempty"`
}

// ReadyDeleted returns a condition that indicates the target object hierarchy
// has successfully been cleaned up, and the InControlPlaneOverride object is
// ready for garbage collection.
func ReadyDeleted() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             "Deleted",
	}
}

// ReadyTraversed returns a condition that indicates whether
// the target object hierarchy has successfully been traversed or not.
func ReadyTraversed(s corev1.ConditionStatus) xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             s,
		LastTransitionTime: metav1.Now(),
		Reason:             "Traversed",
	}
}

var (
	// InControlPlaneOverrideKind is the kind of the InControlPlaneOverride.
	InControlPlaneOverrideKind = reflect.TypeOf(InControlPlaneOverride{}).Name()
)

func init() {
	SchemeBuilder.Register(&InControlPlaneOverride{}, &InControlPlaneOverrideList{})
}
