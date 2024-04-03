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

	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

const (
	// ResourceCredentialsSecretInClusterKubeconfigKey is the key in the
	// connection secret of the ControlPlane that contains the kubeconfig
	// to be used by running pods in the cluster.
	ResourceCredentialsSecretInClusterKubeconfigKey = "kubeconfig-incluster"

	// ConditionMessageAnnotationKey is the key for the message shown in the
	// message column in kubectl.
	ConditionMessageAnnotationKey = "internal.spaces.upbound.io/message"
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

// SourceSpec is the specification about the source of the Control Plane.
type SourceSpec struct {
	// Git is the configuration for a Git repository to pull the Control Plane
	// Source from.
	//
	// +kubebuilder:validation:Required
	Git GitSourceConfig `json:"git"`
}

// GitSourceConfig is the configuration for a Git repository to pull the
// Control Plane Source from.
type GitSourceConfig struct {
	// URL is the URL of the Git repository to pull the Control Plane Source.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	URL string `json:"url"`

	// Ref is the git reference to checkout, which can be a branch, tag, or
	// commit SHA. Default is the main branch.
	//
	// +kubebuilder:default={"branch":"main"}
	Ref *GitReference `json:"ref,omitempty"`

	// Path is the path within the Git repository to pull the Control Plane
	// Source from. The folder it points to must contain a valid ControlPlaneSource
	// manifest. Default is the root of the repository.
	//
	// +kubebuilder:default="/"
	// +kubebuilder:validation:MinLength=1
	Path string `json:"path"`

	// PullInterval is the interval at which the Git repository should be
	// polled for changes. The format is 1h2m3s. Default is 90s. Minimum is 15s.
	//
	// +kubebuilder:default="90s"
	PullInterval *metav1.Duration `json:"pullInterval"`

	// Auth is the authentication configuration to access the Git repository.
	// Default is no authentication.
	// +kubebuilder:default={"type":"None"}
	Auth GitAuthConfig `json:"auth"`
}

// GitReference is a reference to a git branch, tag or commit SHA.
type GitReference struct {
	// TODO(muvaf): Add validation to check that exactly one of these fields
	// is set.

	// Branch is the git branch to check out.
	Branch *string `json:"branch,omitempty"`

	// Tag is the git tag to check out.
	Tag *string `json:"tag,omitempty"`

	// Commit is the git commit SHA to check out.
	Commit *string `json:"commit,omitempty"`
}

// GitAuthConfig is the configuration for authentication to access a Git.
type GitAuthConfig struct {
	// CASecretRef is a reference to a Secret containing CA certificates to use
	// to verify the Git server's certificate. The secret must contain the key
	// "ca.crt" where the content is a CA certificate. The type of the secret
	// can be "Opaque" or "kubernetes.io/tls".
	//
	// +kubebuilder:validation:Optional
	CASecretRef *xpv1.SecretReference `json:"caSecretRef,omitempty"`

	// Type of the authentication to use. Options are: None, Basic
	// (username/password), BearerToken, SSH. Default is None. The corresponding
	// fields must be set for the chosen authentication type.
	//
	// If you are looking to use OAuth tokens with popular servers (e.g.
	// GitHub, Bitbucket, GitLab) you should use BasicAuth instead of
	// BearerToken. These servers use basic HTTP authentication, with the OAuth
	// token as user or password.
	// Check the documentation of your git server for details.
	//
	// +kubebuilder:default="None"
	// +kubebuilder:validation:Enum=None;Basic;BearerToken;SSH
	Type GitAuthType `json:"type"`

	// Basic is the configuration for basic authentication, i.e. username and
	// password.
	Basic *GitBasicAuth `json:"basic,omitempty"`

	// BearerToken is the configuration for bearer token authentication.
	BearerToken *GitBearerTokenAuth `json:"bearerToken,omitempty"`

	// SSH is the configuration for SSH authentication. Note that the URL must
	// use the SSH protocol (e.g. ssh://github.com/org/repo.git).
	SSH *GitSSHAuth `json:"ssh,omitempty"`
}

// GitBasicAuth is the configuration for basic authentication.
type GitBasicAuth struct {
	// SecretRef is a reference to a Secret containing the username and password.
	// The secret must contain the keys "username" and "password".
	//
	// +kubebuilder:validation:Required
	SecretRef xpv1.SecretReference `json:"secretRef"`
}

// GitBearerTokenAuth is the configuration for bearer token authentication.
type GitBearerTokenAuth struct {
	// SecretRef is a reference to a Secret containing the bearer token.
	// The secret must contain the key "bearerToken".
	//
	// +kubebuilder:validation:Required
	SecretRef xpv1.SecretReference `json:"secretRef"`
}

// GitSSHAuth is the configuration for SSH authentication.
type GitSSHAuth struct {
	// SecretRef is a reference to a Secret containing the SSH key and known
	// hosts list.
	// The secret must contain the key "identity" where the content is a private
	// SSH key. Optionally, it can contain the key "knownHosts" where the content
	// is a known hosts file.
	//
	// +kubebuilder:validation:Required
	SecretRef xpv1.SecretReference `json:"secretRef"`
}

// SourceStatus is the status of the pull and apply operations of resources.
type SourceStatus struct {
	// Reference is the git reference that the Control Plane Source is currently
	// checked out to. This could be a branch, tag or commit SHA.
	Reference string `json:"reference,omitempty"`

	// Revision is always the commit SHA that the Control Plane Source is
	// currently checked out to.
	Revision string `json:"revision,omitempty"`
}

// StorageLocation specifies where to store back data.
type StorageLocation struct {
	// Prefix defines the directory within the control plane's storage location where backups are
	// stored or retrieved.
	// +optional
	// +kubebuilder:validation:MinLength=1
	Prefix *string `json:"prefix,omitempty"`

	// AccessMode specifies the access mode of the control plane's backup storage location.
	// Set to ReadOnly when using restoring an existing control plane to another, so
	// that two control planes aren't backing up to the same location.
	// +optional
	// +kubebuilder:default=ReadWrite
	AccessMode *velerov1.BackupStorageLocationAccessMode `json:"accessMode,omitempty"`
}

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
type ControlPlaneSpec struct {
	// [[GATE:EnableGitSource]] THIS IS AN ALPHA FIELD. Do not use it in production.
	// Source points to a Git repository containing a ControlPlaneSource
	// manifest with the desired state of the ControlPlane's configuration.
	Source *SourceSpec `json:"source,omitempty"`

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
	//
	// +optional
	WriteConnectionSecretToReference *SecretReference `json:"writeConnectionSecretToRef,omitempty"`
	// PublishConnectionDetailsTo specifies the connection secret config which
	// contains a name, metadata and a reference to secret store config to
	// which any connection details for this managed resource should be written.
	// Connection details frequently include the endpoint, username,
	// and password required to connect to the managed resource.
	//
	// +optional
	PublishConnectionDetailsTo *xpv1.PublishConnectionDetailsTo `json:"publishConnectionDetailsTo,omitempty"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// ManagementPolicies specify the array of actions Crossplane is allowed to
	// take on the managed and external resources.
	// This field is planned to replace the DeletionPolicy field in a future
	// release. Currently, both could be set independently and non-default
	// values would be honored if the feature flag is enabled. If both are
	// custom, the DeletionPolicy field will be ignored.
	// See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223
	// and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md
	// +optional
	// +kubebuilder:default={"*"}
	ManagementPolicies xpv1.ManagementPolicies `json:"managementPolicies,omitempty"`
	// DeletionPolicy specifies what will happen to the underlying external
	// resource when this managed resource is deleted - either "Delete" or
	// "Orphan" the external resource.
	// This field is planned to be deprecated in favor of the ManagementPolicy
	// field in a future release. Currently, both could be set independently and
	// non-default values would be honored if the feature flag is enabled.
	// See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223
	// +optional
	// +kubebuilder:default=Delete
	DeletionPolicy xpv1.DeletionPolicy `json:"deletionPolicy,omitempty"`

	// Crossplane defines the configuration for Crossplane.
	Crossplane CrossplaneSpec `json:"crossplane,omitempty"`

	// [[GATE:EnableControlPlaneBackup]] THIS IS AN ALPHA FIELD. Do not use it in production.
	// Backup specifies details about the control planes backup configuration.
	// +optional
	Backup *ControlPlaneBackupSpec `json:"backup,omitempty"`

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
	// +kubebuilder:validation:XValidation:rule="(!has(self.apiGroup) || self.apiGroup == 'spaces.upbound.io') && (self.kind == 'Backup' || self.kind == 'BackupSchedule')",message="source must be a reference to a Backup or BackupSchedule (v1alpha1)"
	// +kubebuilder:validation:XValidation:rule="oldSelf == self",message="source is immutable"
	Source TypedLocalObjectReference `json:"source"`

	// FinishedAt is the time at which the control plane was restored, it's not
	// meant to be set by the user, but rather by the system when the control
	// plane is restored.
	FinishedAt *metav1.Time `json:"finishedAt,omitempty"`
}

// ControlPlaneBackupSpec specifies details about the control planes backup configuration.
type ControlPlaneBackupSpec struct {
	// StorageLocation specifies details about the control planes underlying storage location
	// where backups are stored or retrieved.
	// +optional
	StorageLocation *StorageLocation `json:"storageLocation,omitempty"`
}

// A ControlPlaneStatus represents the observed state of a ControlPlane.
type ControlPlaneStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	ControlPlaneID string `json:"controlPlaneID,omitempty"`
	HostClusterID  string `json:"hostClusterID,omitempty"`

	// [[GATE:EnableGitSource]] SourceStatus is the status of the pull and apply operations of resources
	// from the Source.
	SourceStatus *SourceStatus `json:"source,omitempty"`
}

// TypedLocalObjectReference contains enough information to let you locate the
// typed referenced object inside the same namespace.
// +structType=atomic
type TypedLocalObjectReference struct {
	// APIGroup is the group for the resource being referenced.
	// If APIGroup is not specified, the specified Kind must be in the core API group.
	// For any other third-party types, APIGroup is required.
	// +optional
	APIGroup *string `json:"apiGroup,omitempty"`

	// Kind is the type of resource being referenced
	// +kubebuilder:validation:MinLength=1
	Kind string `json:"kind,omitempty"`

	// Name is the name of resource being referenced
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Crossplane",type="string",JSONPath=".spec.crossplane.version"
// +kubebuilder:printcolumn:name="Synced",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=`.metadata.annotations['internal\.spaces\.upbound\.io/message']`
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

// GetDeletionPolicy of this Environment.
func (mg *ControlPlane) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ManagedControlPlane.
func (mg *ControlPlane) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ControlPlane.
func (mg *ControlPlane) GetProviderConfigReference() *xpv1.Reference {
	return nil
}

// GetPublishConnectionDetailsTo of this ControlPlane.
func (mg *ControlPlane) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
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

// SetDeletionPolicy of this ControlPlane.
func (mg *ControlPlane) SetDeletionPolicy(r xpv1.DeletionPolicy) {}

// SetManagementPolicies of this ManagedControlPlane.
func (mg *ControlPlane) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderReference of this ControlPlane.
func (mg *ControlPlane) SetProviderReference(r *xpv1.Reference) {}

// SetProviderConfigReference of this ControlPlane.
func (mg *ControlPlane) SetProviderConfigReference(r *xpv1.Reference) {}

// SetPublishConnectionDetailsTo of this ControlPlane.
func (mg *ControlPlane) SetPublishConnectionDetailsTo(p *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = p
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
)

func init() {
	SchemeBuilder.Register(&ControlPlane{}, &ControlPlaneList{})
}
