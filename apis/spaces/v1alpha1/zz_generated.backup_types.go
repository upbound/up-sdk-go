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

// Generated from spaces/v1beta1/backup_types.go by ../hack/duplicate_api_type.sh. DO NOT EDIT.

package v1alpha1

import (
	"fmt"
	"reflect"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/upbound/up-sdk-go/apis/common"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Retries",type="integer",JSONPath=".status.retries"
// +kubebuilder:printcolumn:name="TTL",type="string",JSONPath=".spec.ttl"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories=spaces

// Backup represents a single backup of a ControlPlane.
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self.controlPlane == oldSelf.controlPlane",message="backup target controlplane can not be changed after creation"
	// +kubebuilder:validation:XValidation:rule="(!has(self.excludedResources) && !has(oldSelf.excludedResources)) || self.excludedResources == oldSelf.excludedResources",message="backup excluded resources can not be changed after creation"
	// +kubebuilder:validation:XValidation:rule="self.configRef == oldSelf.configRef",message="backup config ref can not be changed after creation"
	Spec   BackupSpec   `json:"spec"`
	Status BackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// BackupList contains a list of Backups.
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

// BackupSpec defines a backup over a set of ControlPlanes.
type BackupSpec struct {
	// ControlPlane is the name of the ControlPlane to backup.
	// Requires "backup" permission on the referenced ControlPlane.
	// +kubebuilder:validation:MinLength=1
	ControlPlane string `json:"controlPlane"`

	BackupDefinition `json:",inline"`
}

// BackupDefinition defines all the parameters for a backup.
type BackupDefinition struct {
	// TTL is the time to live for the backup. After this time, the backup
	// will be eligible for garbage collection. If not set, the backup will
	// not be garbage collected.
	// +optional
	TTL *metav1.Duration `json:"ttl,omitempty"`

	// DeletionPolicy is the policy for the backup.
	// +kube:validation:Enum=Orphan;Delete
	// +kubebuilder:default=Orphan
	DeletionPolicy xpv1.DeletionPolicy `json:"deletionPolicy,omitempty"`

	ControlPlaneBackupConfig `json:",inline"`

	// ConfigRef is a reference to the backup configuration.
	// ApiGroup is optional and defaults to "spaces.upbound.io".
	// Kind is required, and the only supported value is "SharedBackupConfig" at
	// the moment.
	// Name is required.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="(!has(self.apiGroup) || self.apiGroup == 'spaces.upbound.io') && self.kind == 'SharedBackupConfig'",message="backup config ref must be a reference to a SharedBackupConfig"
	// +kubebuilder:validation:XValidation:rule="size(self.name) > 0",message="backup config ref must have a name"
	ConfigRef common.TypedLocalObjectReference `json:"configRef"`
}

// ControlPlaneBackupConfig defines the configuration for a controlplane backup.
type ControlPlaneBackupConfig struct {
	// ExcludedResources is a slice of resource names that are not
	// included in the backup. Used to filter the included extra resources.
	// +optional
	// +listType=set
	ExcludedResources []string `json:"excludedResources,omitempty"`
}

// BackupPhase is a string representation of the phase of a backup.
type BackupPhase string

const (
	// BackupPhasePending means the backup has been accepted by the system, but
	// is not yet being processed.
	BackupPhasePending BackupPhase = "Pending"
	// BackupPhaseInProgress means the backup is currently being processed.
	BackupPhaseInProgress BackupPhase = "InProgress"
	// BackupPhaseCompleted means the backup has been completed.
	BackupPhaseCompleted BackupPhase = "Completed"
	// BackupPhaseFailed means the backup has failed.
	BackupPhaseFailed BackupPhase = "Failed"
	// BackupPhaseDeleted means the backup has been deleted from the bucket, at
	// the best of our knowledge.
	BackupPhaseDeleted BackupPhase = "Deleted"
)

// Condition types for backups
const (
	// BackupCompleted indicates that the backup has completed successfully.
	ConditionTypeCompleted xpv1.ConditionType = "Completed"
	// BackupFailed indicates that the backup has failed.
	ConditionTypeFailed xpv1.ConditionType = "Failed"
)

const (
	// AllBackupsCompleted indicates that all backups have completed successfully.
	AllBackupsCompleted xpv1.ConditionReason = "AllBackupsCompleted"
	// AtLeastOneFailed indicates that at least one backup has failed.
	AtLeastOneFailed xpv1.ConditionReason = "AtLeastOneFailed"
)

// SharedBackupCompleted returns a condition indicating that all backups have
// completed successfully.
func SharedBackupCompleted() xpv1.Condition {
	return xpv1.Condition{
		Type:               ConditionTypeCompleted,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             AllBackupsCompleted,
	}
}

// SharedBackupFailed returns a condition indicating that at least one backup
// has failed.
func SharedBackupFailed(err error) xpv1.Condition {
	return xpv1.Condition{
		Type:               ConditionTypeFailed,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             AtLeastOneFailed,
		Message:            err.Error(),
	}
}

// BackupCompleted returns a condition indicating that the backup has completed
// successfully.
func BackupCompleted() xpv1.Condition {
	return xpv1.Condition{
		Type:               ConditionTypeCompleted,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             BackupSuccessReason,
	}
}

const (
	// BackupFailedReason is the reason for a failed backup.
	BackupFailedReason xpv1.ConditionReason = "BackupFailed"
	// BackupRetryReason is the reason for a backup being retried.
	BackupRetryReason xpv1.ConditionReason = "BackupRetry"
	// BackupSuccessReason is the reason for a successful backup.
	BackupSuccessReason xpv1.ConditionReason = "BackupSuccess"
	// BackupControlPlaneNotFoundReason is the reason for a control plane not being found.
	BackupControlPlaneNotFoundReason xpv1.ConditionReason = "TargetControlPlaneNotFound"
)

// BackupFailed returns a condition indicating that the backup has failed.
func BackupFailed(err error) xpv1.Condition {
	return xpv1.Condition{
		Type:               ConditionTypeFailed,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             BackupFailedReason,
		Message:            err.Error(),
	}
}

// BackupRetry returns a condition indicating that the backup is being retried.
func BackupRetry(msg string) xpv1.Condition {
	return xpv1.Condition{
		Type:               ConditionTypeFailed,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             BackupRetryReason,
		Message:            msg,
	}
}

// BackupControlPlaneNotFound returns a condition indicating that the control
// plane was not found.
func BackupControlPlaneNotFound(ns, name string) xpv1.Condition {
	return xpv1.Condition{
		Type:               ConditionTypeFailed,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             BackupControlPlaneNotFoundReason,
		Message:            fmt.Sprintf("target control plane '%s/%s' not found", ns, name),
	}
}

// BackupStatus represents the observed state of a Backup.
type BackupStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// Phase is the current phase of the backup.
	// +kubebuilder:validation:Enum=Pending;InProgress;Failed;Completed;Deleted
	// +kubebuilder:default=Pending
	Phase BackupPhase `json:"phase,omitempty"`

	// Retries is the number of times the backup has been retried.
	Retries int32 `json:"retries,omitempty"`
}

// PreciseLocalObjectReference references by name and uid.
type PreciseLocalObjectReference struct {
	// Name is the name of the referenced object.
	// +optional
	Name string `json:"name,omitempty"`

	// UID is the UID of the referenced object.
	UID types.UID `json:"uid,omitempty"`
}

var (
	// BackupKind is the kind of a Backup.
	BackupKind = reflect.TypeOf(Backup{}).Name()
)

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
