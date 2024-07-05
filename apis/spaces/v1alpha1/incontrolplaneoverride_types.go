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
	"regexp"
	"strings"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
)

var (
	regexFieldNotDeclared = regexp.MustCompile(`failed to create typed patch object.+field not declared in schema`)
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
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

// PropagationMode denotes the traversal direction on an object's hierarchy.
type PropagationMode string

const (
	// PropagateAscend denotes traversal on a target object hierarchy following
	// the metadata.ownerReferences.
	PropagateAscend PropagationMode = "Ascending"
	// PropagateDescend denotes traversal on a target object hierarchy
	// following the spec.resourceRef & spec.resourceRefs reference fields.
	PropagateDescend PropagationMode = "Descending"
	// PropagateNone denotes that no traversal will be done and only the target
	// object will be visited.
	PropagateNone PropagationMode = "None"
)

// Metadata represents the Kube object metadata.
type Metadata struct {
	// Annotations represents the Kube object annotations.
	// Only the following annotations are allowed to be patched:
	// - crossplane.io/paused
	// - spaces.upbound.io/force-reconcile-at
	// +kubebuilder:validation:XValidation:rule="self.all(k, k == 'crossplane.io/paused' || k == 'spaces.upbound.io/force-reconcile-at')",message="Only the crossplane.io/paused and spaces.upbound.io/force-reconcile-at annotations are allowed"
	// +kubebuilder:validation:MinProperties=1
	// +kubebuilder:validation:MaxProperties=2
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Spec represents the spec of a Kube object.
type Spec struct {
	// ManagementPolicies denotes the management policies of a Crossplane
	// managed resource.
	// +optional
	ManagementPolicies xpv1.ManagementPolicies `json:"managementPolicies,omitempty"`
}

// Patch represents a configuration patch which is serialized into JSON to
// obtain the fully specified intent to be used with server side apply.
// +kubebuilder:validation:MinProperties=1
type Patch struct {
	// +optional
	Metadata *Metadata `json:"metadata,omitempty"`

	// +optional
	Spec *Spec `json:"spec,omitempty"`
}

// InControlPlaneOverrideSpec defines a configuration override
// on a target object hierarchy in a target ControlPlane with the
// given name.
type InControlPlaneOverrideSpec struct {
	// ControlPlaneName is the name of the target ControlPlane where
	// the resource configuration overrides will be applied.
	// +kubebuilder:validation:MinLength=1
	ControlPlaneName string `json:"controlPlaneName"`

	Target corev1.TypedObjectReference `json:"target"`

	// +kubebuilder:validation:Enum=None;Ascending;Descending
	// +kubebuilder:default=None
	PropagationMode PropagationMode `json:"propagationMode"`
	Patch           Patch           `json:"patch"`
}

// PatchState denotes the result of the patch operation on the associated
// target object.
type PatchState string

const (
	// PatchStateSuccess denotes that the target object has successfully been
	// patched.
	PatchStateSuccess PatchState = "Success"
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

// ObjectReference represents the state of an applied patch to an object
// in the target hierarchy.
type ObjectReference struct {
	corev1.TypedObjectReference `json:",inline"`

	// Metadata UID of the patch target object.
	// +optional
	UID *types.UID `json:"uid,omitempty"`

	// +kubebuilder:validation:Enum=Success;Skipped;Error
	Status PatchState `json:"status"`

	// +optional
	Reason *PatchStateReason `json:"reason,omitempty"`

	// Message holds an optional detail message detailing the observed state.
	// +optional
	Message *string `json:"message,omitempty"`
}

// String returns a string representation of the ObjectReference.
func (r *ObjectReference) String() string {
	if r == nil {
		return "nil"
	}
	return strings.Join([]string{r.TypedObjectReference.String(), ", ",
		"UID: '", string(ptr.Deref(r.UID, "")), "', ",
		"Status: ", string(r.Status), ", ",
		"Reason: '", string(ptr.Deref(r.Reason, "")), "', ",
		"Message: '", ptr.Deref(r.Message, ""), "'"}, "")
}

// NotFound returns true if the patch operation has failed because the target
// object was not found.
func (r *ObjectReference) NotFound() bool {
	return r.Status == PatchStateError && ptr.Deref(r.Reason, "") == PatchStateReason(metav1.StatusReasonNotFound)
}

// InControlPlaneOverrideStatus defines the status of an InControlPlaneOverride
// object.
type InControlPlaneOverrideStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// +optional
	ObjectRefs []ObjectReference `json:"objectRefs,omitempty"`
}

// PatchSuccess returns an ObjectReference to the patch target object
// indicating the patch has successfully been applied to the object.
func PatchSuccess(u *unstructured.Unstructured) ObjectReference {
	return ObjectReference{
		TypedObjectReference: corev1.TypedObjectReference{
			APIGroup:  ptr.To(u.GetAPIVersion()),
			Kind:      u.GetKind(),
			Name:      u.GetName(),
			Namespace: ptr.To(u.GetNamespace()),
		},
		UID:    ptr.To(u.GetUID()),
		Status: PatchStateSuccess,
	}
}

// PatchFailure returns an ObjectReference to the patch target object together
// with a detail message explaining the transient error encountered.
func PatchFailure(t corev1.TypedObjectReference, uid *types.UID, err error) ObjectReference {
	var reason *PatchStateReason
	sErr := &kerrors.StatusError{}
	ps := PatchStateError
	if errors.As(err, &sErr) {
		reason = ptr.To(PatchStateReason(sErr.Status().Reason))
		switch {
		case kerrors.IsInternalError(err) && regexFieldNotDeclared.MatchString(sErr.Status().Message):
			ps = PatchStateSkipped
			reason = ptr.To(PatchStateReasonSchemaMismatch)

		case sErr.Status().Details != nil:
			for _, c := range sErr.Status().Details.Causes {
				if c.Type == metav1.CauseTypeFieldManagerConflict {
					ps = PatchStateSkipped
					reason = ptr.To(PatchStateReasonConflict)
					break
				}
			}
		}
	}

	return ObjectReference{
		TypedObjectReference: t,
		UID:                  uid,
		Status:               ps,
		Reason:               reason,
		Message:              ptr.To(err.Error()),
	}
}

// Deleted returns a condition that indicates the target object hierarchy
// has successfully been cleaned up, and the InControlPlaneOverride object is
// ready for garbage collection.
func Deleted() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             "Deleted",
	}
}

// Traversed returns a condition that indicates the target object hierarchy has
// successfully been traversed.
func Traversed() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             "Traversed",
	}
}

func init() {
	SchemeBuilder.Register(&InControlPlaneOverride{}, &InControlPlaneOverrideList{})
}
