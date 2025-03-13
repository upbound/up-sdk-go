// Copyright 2025 Upbound Inc
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
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

const (
	// ConditionReasonManifestInvalid says that the ReferencedObject.spec.forProvider is
	// invalid.
	ConditionReasonManifestInvalid xpv1.ConditionReason = "ManifestInvalid"

	// ConditionReasonCompositeNotFound says that the composite was not found.
	// It's used with the Synced condition.
	ConditionReasonCompositeNotFound xpv1.ConditionReason = "CompositeNotFound"

	// ConditionReasonCompositeError says that some other composite error
	// occurred. It's used with the Synced condition.
	ConditionReasonCompositeError xpv1.ConditionReason = "CompositeError"

	// ConditionReasonClaimNotFound says that the claim was not found. It's
	// used with the Synced condition.
	ConditionReasonClaimNotFound xpv1.ConditionReason = "ClaimNotFound"

	// ConditionReasonClaimError says that some other claim error
	// occurred. It's used with the Synced condition.
	ConditionReasonClaimError xpv1.ConditionReason = "ClaimError"

	// ConditionReasonRemoteReferenceNotFound says that the referenced object was not
	// found. It's used with the Synced condition.
	ConditionReasonRemoteReferenceNotFound xpv1.ConditionReason = "RemoteReferencedObjectNotFound"

	// ConditionReasonRemoteJSONPathInvalid says that the JSONPath in the
	// ReferencedObject.spec.composite.JSONPath is invalid.
	ConditionReasonRemoteJSONPathInvalid xpv1.ConditionReason = "JSONPathInvalid"

	// ConditionReasonRemoteReferencedObjectError says that some other remote reference
	// error occurred. It's used with the Synced condition.
	ConditionReasonRemoteReferencedObjectError xpv1.ConditionReason = "RemoteReferencedObjectError"
)

// ManifestInvalid returns a condition that indicates that the
// ReferencedObject.spec.forProvider is invalid.
func ManifestInvalid() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeSynced,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ConditionReasonManifestInvalid,
	}
}

// CompositeError returns a condition that indicates that the composite
// reference has an error.
func CompositeError(err error) xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeSynced,
		Status:             corev1.ConditionFalse,
		Reason:             ConditionReasonCompositeError,
		Message:            err.Error(),
		LastTransitionTime: metav1.Now(),
	}
}

// CompositeNotFound returns a condition that indicates that the composite
// reference.
func CompositeNotFound() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeSynced,
		Status:             corev1.ConditionFalse,
		Reason:             ConditionReasonCompositeNotFound,
		Message:            "Composite not found",
		LastTransitionTime: metav1.Now(),
	}
}

// ClaimReferenceNotFound returns a condition that indicates that the claim
// is undefined at the given JSONPath.
func ClaimReferenceNotFound(gvk schema.GroupVersionKind, nname types.NamespacedName, jsonPath string) xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeSynced,
		Status:             corev1.ConditionFalse,
		Reason:             ConditionReasonClaimError,
		Message:            fmt.Sprintf("Field %s in claim %s %q is undefined", jsonPath, gvk.Kind, nname),
		LastTransitionTime: metav1.Now(),
	}
}

// ClaimError returns a condition that indicates an error with the claim.
func ClaimError(gvk schema.GroupVersionKind, nname types.NamespacedName, err error) xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeSynced,
		Status:             corev1.ConditionFalse,
		Reason:             ConditionReasonClaimError,
		Message:            fmt.Sprintf("Claim %s %q error: %v", gvk.Kind, nname, err),
		LastTransitionTime: metav1.Now(),
	}
}

// ClaimNotFound returns a condition that indicates that the claim was not found.
func ClaimNotFound(gvk schema.GroupVersionKind, nname types.NamespacedName) xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeSynced,
		Status:             corev1.ConditionFalse,
		Reason:             ConditionReasonClaimNotFound,
		Message:            fmt.Sprintf("Claim %s %q not found", gvk.Kind, nname),
		LastTransitionTime: metav1.Now(),
	}
}

// RemoteReferencedObjectNotFound returns a condition that indicates that the
// remote referenced object was not found.
func RemoteReferencedObjectNotFound(gvk schema.GroupVersionKind, nname types.NamespacedName) xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeSynced,
		Status:             corev1.ConditionTrue,
		Reason:             ConditionReasonRemoteReferenceNotFound,
		Message:            fmt.Sprintf("Referenced %s %q not found", gvk.Kind, nname),
		LastTransitionTime: metav1.Now(),
	}
}

// RemoteJSONPathInvalid returns a condition that indicates that the JSONPath in
// the ReferencedObject.spec.composite.JSONPath is invalid.
func RemoteJSONPathInvalid(err error) xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeSynced,
		Status:             corev1.ConditionFalse,
		Reason:             ConditionReasonRemoteJSONPathInvalid,
		Message:            err.Error(),
		LastTransitionTime: metav1.Now(),
	}
}

// RemoteReferencedObjectError returns a condition that indicates that the
// remote referenced object has an error.
func RemoteReferencedObjectError(err error) xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeSynced,
		Status:             corev1.ConditionFalse,
		Reason:             ConditionReasonRemoteReferencedObjectError,
		Message:            err.Error(),
		LastTransitionTime: metav1.Now(),
	}
}

const (
	// ConditionReasonReadyObjectSynced is used when readiness policy is
	// ObjectSynced and the sync was successful.
	ConditionReasonReadyObjectSynced = "ObjectSynced"

	// ConditionReasonReadyObjectNotSynced is used when readiness policy is
	// ObjectExists and the sync was not successful.
	ConditionReasonReadyObjectNotSynced = "ObjectNotSynced"

	// ConditionReasonReadyObjectSyncedUnknown is used when readiness policy is
	// ObjectExists and the sync state is unknown.
	ConditionReasonReadyObjectSyncedUnknown = "ObjectSyncedUnknown"

	// ConditionReasonReadyObjectNotFound is used when the object is not found.
	ConditionReasonReadyObjectNotFound = "ObjectNotFound"

	// ConditionReasonReadyObjectExists is used when the object is found and
	// the readiness policy is ObjectExists.
	ConditionReasonReadyObjectExists = "ObjectExists"

	// ConditionReasonReadyObjectAllConditionsTrue is used when the object's
	// condition are all true and the readiness policy is
	// ObjectConditionsAllTrue.
	ConditionReasonReadyObjectAllConditionsTrue = "ObjectAllConditionsTrue"

	// ConditionReasonReadyObjectNotAllConditionsTrue is used when the object's
	// condition are not all true and the readiness policy is
	// ObjectConditionsAllTrue.
	ConditionReasonReadyObjectNotAllConditionsTrue = "ObjectNotAllConditionsTrue"

	// ConditionReasonReadyObjectInvalid is used when the object is invalid.
	ConditionReasonReadyObjectInvalid = "ObjectInvalid"

	// ConditionReasonReadyNoReadyCondition is used when the object has no ready
	// condition.
	ConditionReasonReadyNoReadyCondition = "NoReadyCondition"
)

// ReadyObjectSynced returns a condition that indicates that the object is
// synced because the readiness policy is ObjectSynced and the sync was
// successful.
func ReadyObjectSynced() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ConditionReasonReadyObjectSynced,
	}
}

// UnreadyObjectNotSynced returns a condition that indicates that the object is
// not synced because the readiness policy is ObjectExists and the sync was not
// successful.
func UnreadyObjectNotSynced() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ConditionReasonReadyObjectNotSynced,
	}
}

// UnreadyObjectSyncedUnknown returns a condition that indicates that the object
// is synced because the readiness policy is ObjectExists and the sync state is
// unknown.
func UnreadyObjectSyncedUnknown() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             corev1.ConditionUnknown,
		LastTransitionTime: metav1.Now(),
		Reason:             ConditionReasonReadyObjectSyncedUnknown,
	}
}

// UnreadyObjectNotFound returns a condition that indicates that the object is
// not found.
func UnreadyObjectNotFound() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ConditionReasonReadyObjectNotFound,
	}
}

// ReadyObjectExists returns a condition that indicates that the object is found
// and the readiness policy is ObjectExists.
func ReadyObjectExists() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ConditionReasonReadyObjectExists,
	}
}

// ReadyObjectAllConditionsTrue returns a condition that indicates that the
// object's conditions are all true and the readiness policy is
// ObjectConditionsAllTrue.
func ReadyObjectAllConditionsTrue() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ConditionReasonReadyObjectAllConditionsTrue,
	}
}

// UnreadyObjectNotAllConditionsTrue returns a condition that indicates that the
// object's conditions are not all true and the readiness policy is
// ObjectConditionsAllTrue.
func UnreadyObjectNotAllConditionsTrue() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ConditionReasonReadyObjectNotAllConditionsTrue,
	}
}

// UnreadyObjectInvalid returns a condition that indicates that the object is
// invalid.
func UnreadyObjectInvalid() xpv1.Condition {
	return xpv1.Condition{
		Type:               xpv1.TypeReady,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ConditionReasonReadyObjectInvalid,
	}
}
