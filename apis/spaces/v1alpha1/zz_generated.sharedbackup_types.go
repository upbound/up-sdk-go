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

// Generated from spaces/v1beta1/sharedbackup_types.go by ../hack/duplicate_api_type.sh. DO NOT EDIT.

package v1alpha1

import (
	"reflect"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SharedBackupLabelKey is the label key used to identify Backups that are
// managed by a SharedBackup.
const SharedBackupLabelKey = "spaces.upbound.io/sharedbackup"

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Completed",type="string",JSONPath=`.metadata.annotations.sharedbackup\.internal\.spaces\.upbound\.io/completed`
// +kubebuilder:printcolumn:name="Failed",type="string",JSONPath=`.metadata.annotations.sharedbackup\.internal\.spaces\.upbound\.io/failed`
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=`.metadata.annotations.sharedbackup\.internal\.spaces\.upbound\.io/provisioned`
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories=spaces

// SharedBackup defines a backup over a set of ControlPlanes.
type SharedBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self.controlPlaneSelector == oldSelf.controlPlaneSelector",message="shared backup ControlPlane selectors can not be changed after creation"
	// +kubebuilder:validation:XValidation:rule="(!has(self.excludedResources) && !has(oldSelf.excludedResources)) || self.excludedResources == oldSelf.excludedResources",message="shared backup excluded resources can not be changed after creation"
	Spec   SharedBackupSpec   `json:"spec"`
	Status SharedBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// SharedBackupList contains a list of SharedBackups.
type SharedBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedBackup `json:"items"`
}

// SharedBackupSpec defines a backup over a set of ControlPlanes.
type SharedBackupSpec struct {
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

	BackupDefinition `json:",inline"`
}

// SharedBackupStatus represents the observed state of a SharedBackup.
type SharedBackupStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// Phase represents the current phase of the SharedBackup.
	// +kubebuilder:validation:Enum=Pending;InProgress;Failed;Completed
	// +kubebuilder:default=Pending
	Phase BackupPhase `json:"phase,omitempty"`

	// SelectedControlPlanes represents the names of the selected ControlPlanes.
	// +optional
	// +listType=set
	SelectedControlPlanes []string `json:"selectedControlPlanes,omitempty"`

	// Failed is the list of ControlPlanes for which the backup failed.
	// +optional
	// +listType=set
	Failed []string `json:"failed,omitempty"`

	// Completed is the list of ControlPlanes for which the backup completed successfully.
	// +optional
	// +listType=set
	Completed []string `json:"completed,omitempty"`
}

var (
	// SharedBackupKind is the kind of a SharedBackup.
	SharedBackupKind = reflect.TypeOf(SharedBackup{}).Name()
)

func init() {
	SchemeBuilder.Register(&SharedBackup{}, &SharedBackupList{})
}
