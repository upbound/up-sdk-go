// Copyright 2025 Upbound Inc
// All rights reserved

// Package v1alpha1 contains pkg.upbound.io APIs in version v1alpha1.
package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	pkgv1 "github.com/crossplane/crossplane/v2/apis/pkg/v1"
)

// TODO(sttts): Move to up-sdk-go

var _ pkgv1.Package = &RemoteConfiguration{}

// +kubebuilder:object:root=true
// +genclient
// +genclient:nonNamespaced

// A RemoteConfiguration installs an OCI compatible configuration package in
// remote mode by creating the claim CustomResourceDefinitions only.
//
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="INSTALLED",type="string",JSONPath=".status.conditions[?(@.type=='Installed')].status"
// +kubebuilder:printcolumn:name="HEALTHY",type="string",JSONPath=".status.conditions[?(@.type=='Healthy')].status"
// +kubebuilder:printcolumn:name="PACKAGE",type="string",JSONPath=".spec.package"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={upbound,pkg}
type RemoteConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RemoteConfigurationSpec   `json:"spec"`
	Status RemoteConfigurationStatus `json:"status,omitempty"`
}

// RemoteConfigurationSpec specifies the configuration of a RemoteConfiguration..
type RemoteConfigurationSpec struct {
	// +kubebuilder:validation:XValidation:rule="self.skipDependencyResolution",message="skipDependencyResolution must be true"
	pkgv1.PackageSpec `json:",inline"`
}

// RemoteConfigurationStatus represents the observed state of a RemoteConfiguration.
type RemoteConfigurationStatus struct {
	xpv1.ConditionedStatus `json:",inline"`
	pkgv1.PackageStatus    `json:",inline"`
}

// +kubebuilder:object:root=true

// RemoteConfigurationList contains a list of RemoteConfiguration.
type RemoteConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RemoteConfiguration `json:"items"`
}

// Implement XP Package interface for RemoteConfiguration.
var _ pkgv1.Package = &RemoteConfiguration{}

// GetAppliedImageConfigRefs returns the applied image config references.
func (in *RemoteConfiguration) GetAppliedImageConfigRefs() []pkgv1.ImageConfigRef {
	return in.Status.AppliedImageConfigRefs
}

// SetAppliedImageConfigRefs sets the applied image config references.
func (in *RemoteConfiguration) SetAppliedImageConfigRefs(refs ...pkgv1.ImageConfigRef) {
	in.Status.AppliedImageConfigRefs = refs
}

// ClearAppliedImageConfigRef clears the applied image config reference for a given reason.
func (in *RemoteConfiguration) ClearAppliedImageConfigRef(reason pkgv1.ImageConfigRefReason) {
	for i := range in.Status.AppliedImageConfigRefs {
		if in.Status.AppliedImageConfigRefs[i].Reason == reason {
			in.Status.AppliedImageConfigRefs = append(in.Status.AppliedImageConfigRefs[:i], in.Status.AppliedImageConfigRefs[i+1:]...)
			return
		}
	}
}

// GetResolvedSource returns the resolved source package.
func (in *RemoteConfiguration) GetResolvedSource() string {
	return in.Status.ResolvedPackage
}

// SetResolvedSource sets the resolved source package.
func (in *RemoteConfiguration) SetResolvedSource(s string) {
	in.Status.ResolvedPackage = s
}

// SetConditions sets the conditions.
func (in *RemoteConfiguration) SetConditions(c ...xpv1.Condition) {
	in.Status.SetConditions(c...)
}

// GetCondition gets the condition.
func (in *RemoteConfiguration) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return in.Status.GetCondition(ct)
}

// CleanConditions removes all conditions.
func (in *RemoteConfiguration) CleanConditions() {
	in.Status.Conditions = []xpv1.Condition{}
}

// GetSource gets the source package.
func (in *RemoteConfiguration) GetSource() string {
	return in.Spec.Package
}

// SetSource sets the source package.
func (in *RemoteConfiguration) SetSource(s string) {
	in.Spec.Package = s
}

// GetActivationPolicy gets the activation policy.
func (in *RemoteConfiguration) GetActivationPolicy() *pkgv1.RevisionActivationPolicy {
	return in.Spec.RevisionActivationPolicy
}

// SetActivationPolicy sets the activation policy.
func (in *RemoteConfiguration) SetActivationPolicy(a *pkgv1.RevisionActivationPolicy) {
	in.Spec.RevisionActivationPolicy = a
}

// GetPackagePullSecrets gets the package pull secrets.
func (in *RemoteConfiguration) GetPackagePullSecrets() []corev1.LocalObjectReference {
	return in.Spec.PackagePullSecrets
}

// SetPackagePullSecrets sets the package pull secrets.
func (in *RemoteConfiguration) SetPackagePullSecrets(s []corev1.LocalObjectReference) {
	in.Spec.PackagePullSecrets = s
}

// GetPackagePullPolicy gets the package pull policy.
func (in *RemoteConfiguration) GetPackagePullPolicy() *corev1.PullPolicy {
	return in.Spec.PackagePullPolicy
}

// SetPackagePullPolicy sets the package pull policy.
func (in *RemoteConfiguration) SetPackagePullPolicy(i *corev1.PullPolicy) {
	in.Spec.PackagePullPolicy = i
}

// GetRevisionHistoryLimit gets the revision history limit.
func (in *RemoteConfiguration) GetRevisionHistoryLimit() *int64 {
	return in.Spec.RevisionHistoryLimit
}

// SetRevisionHistoryLimit sets the revision history limit.
func (in *RemoteConfiguration) SetRevisionHistoryLimit(l *int64) {
	in.Spec.RevisionHistoryLimit = l
}

// GetIgnoreCrossplaneConstraints gets the ignore crossplane constraints.
func (in *RemoteConfiguration) GetIgnoreCrossplaneConstraints() *bool {
	return in.Spec.IgnoreCrossplaneConstraints
}

// SetIgnoreCrossplaneConstraints sets the ignore crossplane constraints.
func (in *RemoteConfiguration) SetIgnoreCrossplaneConstraints(b *bool) {
	in.Spec.IgnoreCrossplaneConstraints = b
}

// GetCurrentRevision gets the current revision.
func (in *RemoteConfiguration) GetCurrentRevision() string {
	return in.Status.CurrentRevision
}

// SetCurrentRevision sets the current revision.
func (in *RemoteConfiguration) SetCurrentRevision(r string) {
	in.Status.CurrentRevision = r
}

// GetCurrentIdentifier gets the current identifier.
func (in *RemoteConfiguration) GetCurrentIdentifier() string {
	return in.Status.CurrentIdentifier
}

// SetCurrentIdentifier sets the current identifier.
func (in *RemoteConfiguration) SetCurrentIdentifier(r string) {
	in.Status.CurrentIdentifier = r
}

// GetSkipDependencyResolution gets the skip dependency resolution.
func (in *RemoteConfiguration) GetSkipDependencyResolution() *bool {
	return in.Spec.SkipDependencyResolution
}

// SetSkipDependencyResolution sets the skip dependency resolution.
func (in *RemoteConfiguration) SetSkipDependencyResolution(skip *bool) {
	in.Spec.SkipDependencyResolution = skip
}

// GetCommonLabels gets the common labels.
func (in *RemoteConfiguration) GetCommonLabels() map[string]string {
	return in.Spec.CommonLabels
}

// SetCommonLabels sets the common labels.
func (in *RemoteConfiguration) SetCommonLabels(l map[string]string) {
	in.Spec.CommonLabels = l
}
