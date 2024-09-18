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

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpaceBackupScheduleLabelKey is the label key used to identify SpaceBackups created by
// a SpaceBackupSchedule.
const SpaceBackupScheduleLabelKey = "spaces.upbound.io/spacebackupschedule"

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Schedule",type="string",JSONPath=".spec.schedule"
// +kubebuilder:printcolumn:name="LastBackup",type="date",JSONPath=".status.lastBackup"
// +kubebuilder:printcolumn:name="TTL",type="string",JSONPath=".spec.ttl"
// +kubebuilder:printcolumn:name="Suspended",type="boolean",JSONPath=".spec.suspend"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories=spaces

// SpaceBackupSchedule represents a single ControlPlane schedule for Backups.
type SpaceBackupSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpaceBackupScheduleSpec   `json:"spec"`
	Status SpaceBackupScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpaceBackupScheduleList contains a list of SpaceBackupSchedules.
type SpaceBackupScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []SpaceBackupSchedule `json:"items"`
}

// SpaceBackupScheduleDefinition defines the schedule for a SpaceBackup.
type SpaceBackupScheduleDefinition struct {
	// Suspend specifies whether the schedule is suspended. If true, no
	// SpaceBackups will be created, but running backups will be allowed to
	// complete.
	// +optional
	Suspend bool `json:"suspend,omitempty"`

	// Schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	// +kubebuilder:validation:MinLength=1
	Schedule string `json:"schedule"`

	SpaceBackupDefinition `json:",inline"`
}

// SpaceBackupScheduleSpec defines a space backup schedule.
type SpaceBackupScheduleSpec struct {
	// UseOwnerReferencesBackup specifies whether an ownership chain should be
	// established between this resource and the Backup it creates.
	// If set to true, the Backup will be garbage collected when this resource
	// is deleted.
	// +optional
	UseOwnerReferencesInBackup bool `json:"useOwnerReferencesInBackup,omitempty"`

	SpaceBackupScheduleDefinition `json:",inline"`
}

// SpaceBackupScheduleStatus represents the observed state of a SpaceBackupSchedule.
type SpaceBackupScheduleStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// LastBackup is the last time a Backup was run for this
	// Schedule schedule
	// +optional
	LastBackup *metav1.Time `json:"lastBackup,omitempty"`
}

var (
	// SpaceBackupScheduleKind is the kind of SpaceBackupSchedule.
	SpaceBackupScheduleKind = reflect.TypeOf(SpaceBackupSchedule{}).Name()
)

func init() {
	SchemeBuilder.Register(&SpaceBackupSchedule{}, &SpaceBackupScheduleList{})
}
