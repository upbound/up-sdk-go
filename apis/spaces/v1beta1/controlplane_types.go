// Copyright 2023 Upbound Inc
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

package v1beta1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

	"github.com/upbound/up-sdk-go/apis/common"
)

const (
	// ResourceCredentialsSecretInClusterKubeconfigKey is the key in the
	// connection secret of the ControlPlane that contains the kubeconfig
	// to be used by running pods in the cluster.
	ResourceCredentialsSecretInClusterKubeconfigKey = "kubeconfig-incluster"

	// ConditionMessageAnnotationKey is the key for the message shown in the
	// message column in kubectl.
	ConditionMessageAnnotationKey = "internal.spaces.upbound.io/message"

	// ControlPlaneGroupLabelKey is the key used to identify namespaces as groups. The
	// value will be "true".
	ControlPlaneGroupLabelKey = "spaces.upbound.io/group"
	// ControlPlaneGroupProtectionKey is the key used to prevent deletion of groups
	// via the Spaces API. Deletion will not be protected if the underlying namespace
	// is deleted.
	ControlPlaneGroupProtectionKey = "spaces.upbound.io/group-deletion-protection"
)

// GitAuthType is the type of authentication to use to access a Git repository.
type GitAuthType string

// GitAuthType constants.
const (
	GitAuthTypeNone        = "None"
	GitAuthTypeBasic       = "Basic"
	GitAuthTypeBearerToken = "BearerToken"
	GitAuthTypeSSH         = "SSH"

	AuthSecretKeyUsername      = "username"
	AuthSecretKeyPassword      = "password"
	AuthSecretKeyBearerToken   = "bearerToken"
	AuthSecretKeySSHIdentity   = "identity"
	AuthSecretKeySSHKnownHosts = "knownHosts"

	CASecretKeyCAFile = "ca.crt"
)

const (
	// KubeCompositionAnnotation is an optional, alpha-level annotation that
	// selects the KubeControlPlane composition for a specific ControlPlane.
	// The default value is "k8s".
	//
	// It is gated by the "EnableKine" feature gate.
	KubeCompositionAnnotation = "internal.spaces.upbound.io/kube-composition"
	// FeaturesAnnotation is an optional annotation that enables features
	// gates within the control plane compositions. Value should be defined
	// as an inline map of key value pairs expressing features to be enabled.
	// For example: `{"featureA": true,"featureB": false}`. The default value is
	// empty (no features enabled).
	FeaturesAnnotation = "internal.spaces.upbound.io/features"
	// TierLimitsAnnotation is an optional annotation that specifies the limits
	// applied to the control plane, as metered by the mxp-account-gate. These
	// limits are only applicable when the account gate is enabled using the
	// features annotation.
	TierLimitsAnnotation = "internal.spaces.upbound.io/tier-limits"
)

// CrossplaneUpgradeChannel is the channel for Crossplane upgrades.
type CrossplaneUpgradeChannel string

const (
	// CrossplaneUpgradeNone disables auto-upgrades and keeps the control plane at its current version of Crossplane.
	CrossplaneUpgradeNone CrossplaneUpgradeChannel = "None"

	// CrossplaneUpgradePatch automatically upgrades the control plane to the latest supported patch version when it
	// becomes available while keeping the minor version the same.
	CrossplaneUpgradePatch CrossplaneUpgradeChannel = "Patch"

	// CrossplaneUpgradeStable automatically upgrades the control plane to the latest supported patch release on minor
	// version N-1, where N is the latest supported minor version.
	CrossplaneUpgradeStable CrossplaneUpgradeChannel = "Stable"

	// CrossplaneUpgradeRapid automatically upgrades the cluster to the latest supported patch release on the latest
	// supported minor version.
	CrossplaneUpgradeRapid CrossplaneUpgradeChannel = "Rapid"
)

// CrossplaneState is the running state for the crossplane and provider workloads.
type CrossplaneState string

const (
	// CrossplaneStateRunning switches the crossplane and provider workloads to
	// the running state by scaling up them.
	CrossplaneStateRunning CrossplaneState = "Running"

	// CrossplaneStatePaused switches the crossplane and provider workloads to
	// the paused state by scaling down them.
	CrossplaneStatePaused CrossplaneState = "Paused"
)

// CrossplaneAutoUpgradeSpec defines the auto upgrade policy for Crossplane.
type CrossplaneAutoUpgradeSpec struct {
	// Channel defines the upgrade channels for Crossplane. We support the following channels where 'Stable' is the
	// default:
	// - None: disables auto-upgrades and keeps the control plane at its current version of Crossplane.
	// - Patch: automatically upgrades the control plane to the latest supported patch version when it
	//   becomes available while keeping the minor version the same.
	// - Stable: automatically upgrades the control plane to the latest supported patch release on minor
	//   version N-1, where N is the latest supported minor version.
	// - Rapid: automatically upgrades the cluster to the latest supported patch release on the latest
	//   supported minor version.
	// +optional
	// +kubebuilder:default="Stable"
	// +kubebuilder:validation:Enum="None";"Patch";"Stable";"Rapid"
	Channel *CrossplaneUpgradeChannel `json:"channel,omitempty"`
}

// CrossplaneSpec defines the configuration for Crossplane.
type CrossplaneSpec struct {
	// Version is the version of Universal Crossplane to install.
	// +optional
	Version *string `json:"version,omitempty"`

	// AutoUpgrades defines the auto upgrade configuration for Crossplane.
	// +optional
	// +kubebuilder:default={"channel":"Stable"}
	AutoUpgradeSpec *CrossplaneAutoUpgradeSpec `json:"autoUpgrade,omitempty"`

	// State defines the state for crossplane and provider workloads. We support
	// the following states where 'Running' is the default:
	// - Running: Starts/Scales up all crossplane and provider workloads in the ControlPlane
	// - Paused: Pauses/Scales down all crossplane and provider workloads in the ControlPlane
	// +optional
	// +kubebuilder:validation:Enum=Running;Paused
	// +kubebuilder:default=Running
	State *CrossplaneState `json:"state,omitempty"`
}

// A SecretReference is a reference to a secret in an arbitrary namespace.
type SecretReference struct {
	// Name of the secret.
	Name string `json:"name"`

	// Namespace of the secret. If omitted, it is equal to
	// the namespace of the resource containing this reference as a field.
	// +optional
	Namespace string `json:"namespace"`
}

// A ControlPlaneSpec represents the desired state of the ControlPlane.
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.restore) || has(self.restore)",message="[[GATE:EnableSharedBackup]] restore source can not be unset"
// +kubebuilder:validation:XValidation:rule="has(oldSelf.restore) || !has(self.restore)",message="[[GATE:EnableSharedBackup]] restore source can not be set after creation"
// +kubebuilder:validation:XValidation:rule="!has(self.crossplane.autoUpgrade) || self.crossplane.autoUpgrade.channel != \"None\" || self.crossplane.version != \"\"",message="\"version\" cannot be empty when upgrade channel is \"None\""
type ControlPlaneSpec struct {
	// WriteConnectionSecretToReference specifies the namespace and name of a
	// Secret to which any connection details for this managed resource should
	// be written. Connection details frequently include the endpoint, username,
	// and password required to connect to the managed resource.
	// This field is planned to be replaced in a future release in favor of
	// PublishConnectionDetailsTo. Currently, both could be set independently
	// and connection details would be published to both without affecting
	// each other.
	//
	// If omitted, it is defaulted to the namespace of the ControlPlane.
	// Deprecated: Use Hub or Upbound identities instead.
	// +optional
	WriteConnectionSecretToReference *SecretReference `json:"writeConnectionSecretToRef,omitempty"`

	// Crossplane defines the configuration for Crossplane.
	Crossplane CrossplaneSpec `json:"crossplane,omitempty"`

	// [[GATE:EnableSharedBackup]] THIS IS AN ALPHA FIELD. Do not use it in production.
	// Restore specifies details about the control planes restore configuration.
	// +optional
	// +kubebuilder:validation:XValidation:rule="!has(oldSelf.finishedAt) || oldSelf.finishedAt == self.finishedAt",message="finishedAt is immutable once set"
	Restore *Restore `json:"restore,omitempty"`
}

// Restore specifies details about the backup to restore from.
type Restore struct {
	// Source of the Backup or BackupSchedule to restore from.
	// Require "restore" permission on the referenced Backup or BackupSchedule.
	// ApiGroup is optional and defaults to "spaces.upbound.io".
	// Kind is required, and the only supported kinds are Backup and
	// BackupSchedule at the moment.
	// Name is required.
	// +kubebuilder:validation:XValidation:rule="(!has(self.apiGroup) || self.apiGroup == 'spaces.upbound.io') && (self.kind == 'Backup' || self.kind == 'BackupSchedule')",message="source must be a reference to a Backup or BackupSchedule (v1alpha1)"
	// +kubebuilder:validation:XValidation:rule="oldSelf == self",message="source is immutable"
	Source common.TypedLocalObjectReference `json:"source"`

	// FinishedAt is the time at which the control plane was restored, it's not
	// meant to be set by the user, but rather by the system when the control
	// plane is restored.
	FinishedAt *metav1.Time `json:"finishedAt,omitempty"`
}

// A ControlPlaneStatus represents the observed state of a ControlPlane.
type ControlPlaneStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	// Message is a human-readable message indicating details about why the
	// ControlPlane is in this condition.
	Message        string `json:"message,omitempty"`
	ControlPlaneID string `json:"controlPlaneID,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Crossplane",type="string",JSONPath=".spec.crossplane.version"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=`.status.message`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories=spaces,shortName=ctp;ctps

// ControlPlane defines a managed Crossplane instance.
type ControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ControlPlaneSpec   `json:"spec"`
	Status ControlPlaneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ControlPlaneList contains a list of ControlPlane
type ControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControlPlane `json:"items"`
}

// GetCondition of this ControlPlane.
func (mg *ControlPlane) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetWriteConnectionSecretToReference of this ControlPlane.
func (mg *ControlPlane) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	if mg.Spec.WriteConnectionSecretToReference == nil {
		return nil
	}
	return &xpv1.SecretReference{
		Name:      mg.Spec.WriteConnectionSecretToReference.Name,
		Namespace: mg.Spec.WriteConnectionSecretToReference.Namespace,
	}
}

// SetConditions of this ControlPlane.
func (mg *ControlPlane) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetWriteConnectionSecretToReference of this ControlPlane.
func (mg *ControlPlane) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = &SecretReference{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

// ManagedControlPlane type metadata.
var (
	// ControlPlaneKind is the kind of the ControlPlane.
	ControlPlaneKind = reflect.TypeOf(ControlPlane{}).Name()
	// ControlPlaneListKind is the kind of a list of ControlPlane.
	ControlPlaneListKind = reflect.TypeOf(ControlPlaneList{}).Name()
)

func init() {
	SchemeBuilder.Register(&ControlPlane{}, &ControlPlaneList{})
}
