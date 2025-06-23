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
	"github.com/upbound/up-sdk-go/apis/common"
	spacesv1alpha1 "github.com/upbound/up-sdk-go/apis/spaces/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Retries",type="integer",JSONPath=".status.retries"
// +kubebuilder:printcolumn:name="Failed",type="string",JSONPath=`.metadata.annotations.spacebackup\.internal\.spaces\.upbound\.io/failed`
// +kubebuilder:printcolumn:name="TTL",type="string",JSONPath=".spec.ttl"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories=spaces

// SpaceBackup represents a backup of a Space.
type SpaceBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpaceBackupSpec   `json:"spec"`
	Status SpaceBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpaceBackupList contains a list of SpaceBackups.
type SpaceBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpaceBackup `json:"items"`
}

// SpaceBackupSpec defines a backup over a set of Match
// +kubebuilder:validation:XValidation:rule="!has(self.configRef) && !has(oldSelf.configRef) || (self.configRef == oldSelf.configRef) ",message="spec.configRef can't be changed or set after creation"
// +kubebuilder:validation:XValidation:rule="!has(self.match) && !has(oldSelf.match) || (self.match == oldSelf.match) ",message="spec.match can't be changed or set after creation"
// +kubebuilder:validation:XValidation:rule="!has(self.exclude) && !has(oldSelf.exclude) || (self.exclude == oldSelf.exclude) ",message="spec.exclude can't be changed or set after creation"
// +kubebuilder:validation:XValidation:rule="!has(self.controlPlaneBackups) && !has(oldSelf.controlPlaneBackups) || (self.controlPlaneBackups == oldSelf.controlPlaneBackups) ",message="spec.controlPlaneBackups can't be changed or set after creation"
type SpaceBackupSpec struct {
	SpaceBackupDefinition `json:",inline"`
}

// SpaceBackupDefinition defines all the parameters for a space backup.
type SpaceBackupDefinition struct {
	// ConfigRef is a reference to the space backup configuration.
	// ApiGroup is optional and defaults to "spaces.upbound.io".
	// Kind is required, and the only supported value is "SpaceBackupConfig" at
	// the moment.
	// Name is required.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="(!has(self.apiGroup) || self.apiGroup == 'admin.spaces.upbound.io') && self.kind == 'SpaceBackupConfig'",message="backup config ref must be a reference to a SpaceBackupConfig"
	// +kubebuilder:validation:XValidation:rule="size(self.name) > 0",message="backup config ref must have a name"
	ConfigRef common.TypedLocalObjectReference `json:"configRef"`

	// TTL is the time to live for the backup. After this time, the backup
	// will be eligible for garbage collection. If not set, the backup will
	// not be garbage collected.
	// +optional
	TTL *metav1.Duration `json:"ttl,omitempty"`

	// DeletionPolicy is the policy for the backup.
	// +kube:validation:Enum=Orphan;Delete
	// +kubebuilder:default=Orphan
	DeletionPolicy xpv1.DeletionPolicy `json:"deletionPolicy,omitempty"`

	// Match is the selector for resources that should be included in the backup.
	// By default, we'll back up all Groups and for each Group:
	// - All ControlPlanes.
	// - All Secrets.
	// - All other Space API resources, e.g. SharedBackupConfigs, SharedUpboundPolicies, Backups, etc...
	// +optional
	Match *SpaceBackupResourceSelector `json:"match,omitempty"`

	// Exclude is the selector for resources that should be excluded from the backup.
	// If both Match and Exclude are specified, the Exclude selector will be applied
	// after the Match selector.
	// By default, only SpaceBackups are excluded.
	// +optional
	Exclude *SpaceBackupResourceSelector `json:"exclude,omitempty"`

	// ControlPlaneBackups is the definition of the control plane backups,
	// +optional
	ControlPlaneBackups *spacesv1alpha1.ControlPlaneBackupConfig `json:"controlPlaneBackups,omitempty"`

	// Failures defines the failure tolerance for the backup.
	Failures SpaceBackupFailuresConfig `json:"failures,omitempty"`
}

// SpaceBackupResourceSelector represents a selector for Groups and ControlPlanes.
// An object is going to be matched if any of the provided group selectors
// matches object's group AND any of provided control plane selectors
// matches.
type SpaceBackupResourceSelector struct {
	// Groups specifies the groups selected.
	// A group is matched if any of the group selectors matches, if not specified
	// any group is matched. Group selector is ANDed with all other selectors, so no resource in
	// a group not matching the group selector will be included in the backup.
	// +optional
	Groups *spacesv1alpha1.ResourceSelector `json:"groups,omitempty"`

	// ControlPlanes specifies the control planes selected.
	// A control plane is matched if any of the control plane selectors matches, if not specified
	// any control plane in the selected groups is matched.
	// +optional
	ControlPlanes *spacesv1alpha1.ResourceSelector `json:"controlPlanes,omitempty"`

	// Spaces specifies the spaces selected.
	// +optional
	Secrets *spacesv1alpha1.ResourceSelector `json:"secrets,omitempty"`

	// Extras specifies the extra resources selected.
	// +optional
	Extras []GenericSpaceBackupResourceSelector `json:"extras,omitempty"`
}

// GenericSpaceBackupResourceSelector represents a generic resource selector.
type GenericSpaceBackupResourceSelector struct {
	// APIGroup is the group of the resource.
	// +kubebuilder:validation:Required
	APIGroup string `json:"apiGroup,omitempty"`
	// Kind is the kind of the resource.
	// +kubebuilder:validation:Required
	Kind                            string `json:"kind,omitempty"`
	spacesv1alpha1.ResourceSelector `json:",inline"`
}

// SpaceBackupFailuresConfig defines the failure tolerance for a SpaceBackup.
type SpaceBackupFailuresConfig struct {
	// ControlPlanes is the percentage of control planes that are allowed to fail and still consider the backup successful.
	// Can be specified as an integer (e.g., 50) or a percentage string (e.g., "50%").
	ControlPlanes *intstr.IntOrString `json:"controlPlanes,omitempty"`
}

// SpaceBackupStatus represents the observed state of a SpaceBackup.
type SpaceBackupStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// Phase is the current phase of the backup.
	// +kubebuilder:validation:Enum=Pending;InProgress;Failed;Completed;Deleted
	// +kubebuilder:default=Pending
	Phase spacesv1alpha1.BackupPhase `json:"phase,omitempty"`

	// Retries is the number of times the backup has been retried.
	// +optional
	Retries int32 `json:"retries,omitempty"`

	// ControlPlanes contains details about the control planes that were backed up,
	// total number of control planes and failures.
	ControlPlanes SpaceBackupControlPlanesStatus `json:"controlPlanes,omitempty"`
}

// SpaceBackupControlPlanesStatus contains details about the control planes that were backed up.
type SpaceBackupControlPlanesStatus struct {
	// Total is the total number of control planes that were attempted to be backed up.
	// +optional
	Total int32 `json:"total,omitempty"`
	// Failed is the number of control planes that failed to backup.
	// +optional
	Failed int32 `json:"failed,omitempty"`
}

var (
	// SpaceBackupKind is the kind of a SpaceBackup.
	SpaceBackupKind = reflect.TypeOf(SpaceBackup{}).Name()
	// SpaceBackupListKind is the kind of a SpaceBackupList.
	SpaceBackupListKind = reflect.TypeOf(SpaceBackupList{}).Name()
)

func init() {
	SchemeBuilder.Register(&SpaceBackup{}, &SpaceBackupList{})
}
