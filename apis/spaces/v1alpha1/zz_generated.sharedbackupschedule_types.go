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

// Generated from spaces/v1beta1/sharedbackupschedule_types.go by ../hack/duplicate_api_type.sh. DO NOT EDIT.

package v1alpha1

import (
	"reflect"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SharedBackupScheduleLabelKey is the label key used to identify
// BackupSchedules that are managed by a SharedBackupSchedule.
const SharedBackupScheduleLabelKey = "spaces.upbound.io/sharedbackupschedule"

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Schedule",type="string",JSONPath=".spec.schedule"
// +kubebuilder:printcolumn:name="Suspended",type="boolean",JSONPath=".spec.suspend"
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=`.metadata.annotations.sharedbackupschedule\.internal\.spaces\.upbound\.io/provisioned-total`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories=spaces

// SharedBackupSchedule defines a schedule for SharedBackup on a set of
// ControlPlanes.
type SharedBackupSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SharedBackupScheduleSpec   `json:"spec"`
	Status SharedBackupScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// SharedBackupScheduleList contains a list of SharedBackupSchedules.
type SharedBackupScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is the list of SharedBackupSchedules.
	Items []SharedBackupSchedule `json:"items"`
}

// SharedBackupScheduleSpec defines the desired state of a SharedBackupSchedule.
type SharedBackupScheduleSpec struct {
	// ControlPlaneSelector defines the selector for ControlPlanes to backup.
	// Requires "backup" permission on all ControlPlanes in the same namespace.
	// +kubebuilder:validation:XValidation:rule="(has(self.labelSelectors) || has(self.names)) && (size(self.labelSelectors) > 0 || size(self.names) > 0)",message="either names or a labelSelector must be specified"
	ControlPlaneSelector ResourceSelector `json:"controlPlaneSelector"`

	// UseOwnerReferencesBackup specifies whether an ownership chain should be
	// established between this resource and the Backup it creates.
	// If set to true, the Backup will be garbage collected when this resource
	// is deleted.
	// +optional
	UseOwnerReferencesInBackup bool `json:"useOwnerReferencesInBackup,omitempty"`

	BackupScheduleDefinition `json:",inline"`
}

// SharedBackupScheduleStatus represents the observed state of a SharedBackupSchedule.
type SharedBackupScheduleStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// SelectedControlPlanes is the list of ControlPlanes that are selected
	// for backup.
	// +optional
	// +listType=set
	SelectedControlPlanes []string `json:"selectedControlPlanes,omitempty"`
}

var (
	// SharedBackupScheduleKind is the kind of SharedBackupSchedule.
	SharedBackupScheduleKind = reflect.TypeOf(SharedBackupSchedule{}).Name()
)

func init() {
	SchemeBuilder.Register(&SharedBackupSchedule{}, &SharedBackupScheduleList{})
}
