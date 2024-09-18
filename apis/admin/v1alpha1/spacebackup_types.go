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
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Retries",type="integer",JSONPath=".status.retries"
// +kubebuilder:printcolumn:name="TTL",type="string",JSONPath=".spec.ttl"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories=spaces

// SpaceBackup represents a single backup of a ControlPlane.
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

// SpaceBackupSpec defines a backup over a set of Match.
type SpaceBackupSpec struct {
	SpaceBackupDefinition `json:",inline"`
}

// SpaceBackupDefinition defines all the parameters for a space backup.
type SpaceBackupDefinition struct {
	// ConfigRef is a reference to the backup configuration.
	// ApiGroup is optional and defaults to "spaces.upbound.io".
	// Kind is required, and the only supported value is "SharedBackupConfig" at
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
	Match SpaceBackupResourceSelector `json:"match,omitempty"`

	// Exclude is the selector for resources that should be excluded from the backup.
	// If both Match and Exclude are specified, the Exclude selector will be applied
	// after the Match selector.
	Exclude SpaceBackupResourceSelector `json:"exclude,omitempty"`

	// ControlPlaneBackups is the definition of the control plane backups,
	// +kubebuilder:validation:XValidation:rule="(!has(self.excludedResources) && !has(oldSelf.excludedResources)) || self.excludedResources == oldSelf.excludedResources",message="backup excluded resources can not be changed after creation"
	ControlPlaneBackups spacesv1alpha1.ControlPlaneBackupConfig `json:"controlPlaneBackups,omitempty"`
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
	// APIVersion is the API version of the resource.
	APIVersion string `json:"apiVersion,omitempty"`
	// Kind is the kind of the resource.
	Kind string `json:"kind,omitempty"`
	// Namespaces is the namespaces of the resource.
	spacesv1alpha1.ResourceSelector `json:",inline"`
}

// SpaceBackupStatus represents the observed state of a Backup.
type SpaceBackupStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// Phase is the current phase of the backup.
	// +kubebuilder:validation:Enum=Pending;InProgress;Failed;Completed;Deleted
	// +kubebuilder:default=Pending
	Phase spacesv1alpha1.BackupPhase `json:"phase,omitempty"`

	// Retries is the number of times the backup has been retried.
	Retries int32 `json:"retries,omitempty"`
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
