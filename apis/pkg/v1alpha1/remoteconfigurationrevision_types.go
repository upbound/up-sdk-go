// Copyright 2025 Upbound Inc
// All rights reserved

// Package v1alpha1 contains pkg.upbound.io APIs in version v1alpha1.
package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	pkgv1 "github.com/crossplane/crossplane/apis/pkg/v1"
)

var _ pkgv1.PackageRevision = &RemoteConfigurationRevision{}

// +kubebuilder:object:root=true
// +genclient
// +genclient:nonNamespaced

// An RemoteConfigurationRevision represents a revision of an RemoteConfiguration.
//
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="HEALTHY",type="string",JSONPath=".status.conditions[?(@.type=='Healthy')].status"
// +kubebuilder:printcolumn:name="REVISION",type="string",JSONPath=".spec.revision"
// +kubebuilder:printcolumn:name="IMAGE",type="string",JSONPath=".spec.image"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".spec.desiredState"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={upbound,pkgrev}
type RemoteConfigurationRevision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RemoteConfigurationRevisionSpec   `json:"spec"`
	Status RemoteConfigurationRevisionStatus `json:"status,omitempty"`
}

// RemoteConfigurationRevisionSpec specifies the configuration of an RemoteConfigurationRevision.
type RemoteConfigurationRevisionSpec struct {
	// +kubebuilder:validation:XValidation:rule="self.skipDependencyResolution",message="skipDependencyResolution must be true"
	pkgv1.PackageRevisionSpec `json:",inline"`
}

// RemoteConfigurationRevisionStatus represents the observed state of an RemoteConfigurationRevision.
type RemoteConfigurationRevisionStatus struct {
	pkgv1.PackageRevisionStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// RemoteConfigurationRevisionList contains a list of RemoteConfigurationRevision.
type RemoteConfigurationRevisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RemoteConfigurationRevision `json:"items"`
}

// Implement XP Package interface for RemoteConfigurationRevision.
var _ pkgv1.PackageRevision = &RemoteConfigurationRevision{}

// GetAppliedImageConfigRefs returns the applied image config references.
func (in *RemoteConfigurationRevision) GetAppliedImageConfigRefs() []pkgv1.ImageConfigRef {
	return in.Status.AppliedImageConfigRefs
}

// SetAppliedImageConfigRefs sets the applied image config references.
func (in *RemoteConfigurationRevision) SetAppliedImageConfigRefs(refs ...pkgv1.ImageConfigRef) {
	in.Status.AppliedImageConfigRefs = refs
}

// ClearAppliedImageConfigRef clears the applied image config reference for a given reason.
func (in *RemoteConfigurationRevision) ClearAppliedImageConfigRef(reason pkgv1.ImageConfigRefReason) {
	for i := range in.Status.AppliedImageConfigRefs {
		if in.Status.AppliedImageConfigRefs[i].Reason == reason {
			in.Status.AppliedImageConfigRefs = append(in.Status.AppliedImageConfigRefs[:i], in.Status.AppliedImageConfigRefs[i+1:]...)
			return
		}
	}
}

// GetResolvedSource gets the resolved source package.
func (in *RemoteConfigurationRevision) GetResolvedSource() string {
	return in.Status.ResolvedPackage
}

// SetResolvedSource sets the resolved source package.
func (in *RemoteConfigurationRevision) SetResolvedSource(s string) {
	in.Status.ResolvedPackage = s
}

// SetConditions sets the conditions.
func (in *RemoteConfigurationRevision) SetConditions(c ...xpv1.Condition) {
	in.Status.SetConditions(c...)
}

// GetCondition gets the condition.
func (in *RemoteConfigurationRevision) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return in.Status.GetCondition(ct)
}

// CleanConditions removes all conditions.
func (in *RemoteConfigurationRevision) CleanConditions() {
	in.Status.Conditions = []xpv1.Condition{}
}

// GetObjects returns the objects.
func (in *RemoteConfigurationRevision) GetObjects() []xpv1.TypedReference {
	return in.Status.ObjectRefs
}

// SetObjects sets the objects.
func (in *RemoteConfigurationRevision) SetObjects(c []xpv1.TypedReference) {
	in.Status.ObjectRefs = c
}

// GetSource gets the source package.
func (in *RemoteConfigurationRevision) GetSource() string {
	return in.Spec.Package
}

// SetSource sets the source package.
func (in *RemoteConfigurationRevision) SetSource(s string) {
	in.Spec.Package = s
}

// GetPackagePullSecrets gets the package pull secrets.
func (in *RemoteConfigurationRevision) GetPackagePullSecrets() []corev1.LocalObjectReference {
	return in.Spec.PackagePullSecrets
}

// SetPackagePullSecrets sets the package pull secrets.
func (in *RemoteConfigurationRevision) SetPackagePullSecrets(s []corev1.LocalObjectReference) {
	in.Spec.PackagePullSecrets = s
}

// GetPackagePullPolicy gets the package pull policy.
func (in *RemoteConfigurationRevision) GetPackagePullPolicy() *corev1.PullPolicy {
	return in.Spec.PackagePullPolicy
}

// SetPackagePullPolicy sets the package pull policy.
func (in *RemoteConfigurationRevision) SetPackagePullPolicy(i *corev1.PullPolicy) {
	in.Spec.PackagePullPolicy = i
}

// GetDesiredState gets the desired state.
func (in *RemoteConfigurationRevision) GetDesiredState() pkgv1.PackageRevisionDesiredState {
	return in.Spec.DesiredState
}

// SetDesiredState sets the desired state.
func (in *RemoteConfigurationRevision) SetDesiredState(d pkgv1.PackageRevisionDesiredState) {
	in.Spec.DesiredState = d
}

// GetIgnoreCrossplaneConstraints gets the ignore crossplane constraints.
func (in *RemoteConfigurationRevision) GetIgnoreCrossplaneConstraints() *bool {
	return in.Spec.IgnoreCrossplaneConstraints
}

// SetIgnoreCrossplaneConstraints sets the ignore crossplane constraints.
func (in *RemoteConfigurationRevision) SetIgnoreCrossplaneConstraints(b *bool) {
	in.Spec.IgnoreCrossplaneConstraints = b
}

// GetRevision gets the revision.
func (in *RemoteConfigurationRevision) GetRevision() int64 {
	return in.Spec.Revision
}

// SetRevision sets the revision.
func (in *RemoteConfigurationRevision) SetRevision(r int64) {
	in.Spec.Revision = r
}

// GetSkipDependencyResolution gets the skip dependency resolution.
func (in *RemoteConfigurationRevision) GetSkipDependencyResolution() *bool {
	return in.Spec.SkipDependencyResolution
}

// SetSkipDependencyResolution sets the skip dependency resolution.
func (in *RemoteConfigurationRevision) SetSkipDependencyResolution(skip *bool) {
	in.Spec.SkipDependencyResolution = skip
}

// GetDependencyStatus returns the dependency status.
func (in *RemoteConfigurationRevision) GetDependencyStatus() (found, installed, invalid int64) {
	return in.Status.FoundDependencies, in.Status.InstalledDependencies, in.Status.InvalidDependencies
}

// SetDependencyStatus sets the dependency status.
func (in *RemoteConfigurationRevision) SetDependencyStatus(found, installed, invalid int64) {
	in.Status.FoundDependencies = found
	in.Status.InstalledDependencies = installed
	in.Status.InvalidDependencies = invalid
}

// GetCommonLabels returns the common labels.
func (in *RemoteConfigurationRevision) GetCommonLabels() map[string]string {
	return in.Spec.CommonLabels
}

// SetCommonLabels sets the common labels.
func (in *RemoteConfigurationRevision) SetCommonLabels(l map[string]string) {
	in.Spec.CommonLabels = l
}

// Implement XP Revision List interface for RemoteConfigurationRevisionList.
var _ pkgv1.PackageRevisionList = &RemoteConfigurationRevisionList{}

// GetRevisions returns the revisions.
func (in *RemoteConfigurationRevisionList) GetRevisions() []pkgv1.PackageRevision {
	prs := make([]pkgv1.PackageRevision, len(in.Items))
	for i, r := range in.Items {
		prs[i] = &r
	}
	return prs
}
