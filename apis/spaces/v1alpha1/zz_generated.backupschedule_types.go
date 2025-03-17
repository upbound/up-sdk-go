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

// Generated from spaces/v1beta1/backupschedule_types.go by ../hack/duplicate_api_type.sh. DO NOT EDIT.

package v1alpha1

import (
	"reflect"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupScheduleLabelKey is the label key used to identify Backups created by
// a BackupSchedule.
const BackupScheduleLabelKey = "spaces.upbound.io/backupschedule"

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Schedule",type="string",JSONPath=".spec.schedule"
// +kubebuilder:printcolumn:name="LastBackup",type="date",JSONPath=".status.lastBackup"
// +kubebuilder:printcolumn:name="TTL",type="string",JSONPath=".spec.ttl"
// +kubebuilder:printcolumn:name="ControlPlane",type="string",JSONPath=".spec.controlPlane"
// +kubebuilder:printcolumn:name="Suspended",type="boolean",JSONPath=".spec.suspend"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories=spaces

// BackupSchedule represents a single ControlPlane schedule for Backups.
type BackupSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupScheduleSpec   `json:"spec"`
	Status BackupScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// BackupScheduleList contains a list of BackupSchedules.
type BackupScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []BackupSchedule `json:"items"`
}

// BackupScheduleDefinition defines the schedule for a Backup.
type BackupScheduleDefinition struct {
	// Suspend specifies whether the schedule is suspended. If true, no
	// Backups will be created, but running backups will be allowed to
	// complete.
	// +optional
	Suspend bool `json:"suspend,omitempty"`

	// Schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	// +kubebuilder:validation:MinLength=1
	Schedule string `json:"schedule"`

	BackupDefinition `json:",inline"`
}

// BackupScheduleSpec defines a backup schedule over a set of ControlPlanes.
type BackupScheduleSpec struct {
	// ControlPlane is the name of the ControlPlane to which the schedule
	// applies.
	// Requires "get" permission on the referenced ControlPlane.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="target can not be changed after creation"
	ControlPlane string `json:"controlPlane"`

	// UseOwnerReferencesBackup specifies whether an ownership chain should be
	// established between this resource and the Backup it creates.
	// If set to true, the Backup will be garbage collected when this resource
	// is deleted.
	// +optional
	UseOwnerReferencesInBackup bool `json:"useOwnerReferencesInBackup,omitempty"`

	BackupScheduleDefinition `json:",inline"`
}

// BackupScheduleStatus represents the observed state of a BackupSchedule.
type BackupScheduleStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// LastBackup is the last time a Backup was run for this
	// Schedule schedule
	// +optional
	LastBackup *metav1.Time `json:"lastBackup,omitempty"`
}

var (
	// BackupScheduleKind is the kind of BackupSchedule.
	BackupScheduleKind = reflect.TypeOf(BackupSchedule{}).Name()
)

func init() {
	SchemeBuilder.Register(&BackupSchedule{}, &BackupScheduleList{})
}
