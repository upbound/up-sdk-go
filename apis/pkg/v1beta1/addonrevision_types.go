// Copyright 2025 Upbound Inc
// All rights reserved

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	pkgv1 "github.com/crossplane/crossplane/apis/pkg/v1"
)

// AddOnPackagingType defines an enum for the addon package type.
type AddOnPackagingType string

const (
	// AddOnPackagingTypeHelm represents an addon that is packaged as a
	// Helm chart.
	AddOnPackagingTypeHelm AddOnPackagingType = "Helm"
)

// +kubebuilder:object:root=true
// +genclient
// +genclient:nonNamespaced

// An AddOnRevision represents a revision of an AddOn.
//
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="HEALTHY",type="string",JSONPath=".status.conditions[?(@.type=='Healthy')].status"
// +kubebuilder:printcolumn:name="REVISION",type="string",JSONPath=".spec.revision"
// +kubebuilder:printcolumn:name="IMAGE",type="string",JSONPath=".spec.image"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".spec.desiredState"
// +kubebuilder:printcolumn:name="DEP-FOUND",type="string",JSONPath=".status.foundDependencies"
// +kubebuilder:printcolumn:name="DEP-INSTALLED",type="string",JSONPath=".status.installedDependencies"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={upbound,pkgrev}
type AddOnRevision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:XValidation:rule="!has(oldSelf.helm) || self.helm == oldSelf.helm",message="helm specification is immutable"
	Spec   AddOnRevisionSpec   `json:"spec"`
	Status AddOnRevisionStatus `json:"status,omitempty"`
}

// AddOnRevisionSpec specifies the configuration of an AddOnRevision.
type AddOnRevisionSpec struct {
	pkgv1.PackageRevisionSpec  `json:",inline"`
	PackageRevisionRuntimeSpec `json:",inline"`

	// Helm specific configuration for an addon revision. This field is
	// managed by the addon and should not be modified directly.
	Helm *HelmSpec `json:"helm,omitempty"`
}

// HelmSpec defines the Helm-specific configuration for an addon revision.
type HelmSpec struct {
	// ReleaseName is the name of the Helm release.
	ReleaseName string `json:"releaseName"`
	// ReleaseNamespace is the namespace of the Helm release.
	ReleaseNamespace string `json:"releaseNamespace"`
	// CRDRefs is a list of CRDs that deployed by this addon.
	CRDRefs []string `json:"crdRefs,omitempty"`
}

// AddOnRevisionStatus represents the observed state of an AddOnRevision.
type AddOnRevisionStatus struct {
	pkgv1.PackageRevisionStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// AddOnRevisionList contains a list of AddOnRevision.
type AddOnRevisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddOnRevision `json:"items"`
}

// PackageRevisionRuntimeSpec specifies configuration for the runtime of a
// package revision.
type PackageRevisionRuntimeSpec struct {
	PackageRuntimeSpec `json:",inline"`
}

// Implement XP Package interface for AddOnRevision.
var _ pkgv1.PackageRevision = &AddOnRevision{}

// GetAppliedImageConfigRefs returns the applied image config references.
func (in *AddOnRevision) GetAppliedImageConfigRefs() []pkgv1.ImageConfigRef {
	return in.Status.AppliedImageConfigRefs
}

// SetAppliedImageConfigRefs sets the applied image config references.
func (in *AddOnRevision) SetAppliedImageConfigRefs(refs ...pkgv1.ImageConfigRef) {
	in.Status.AppliedImageConfigRefs = refs
}

// ClearAppliedImageConfigRef clears the applied image config reference for a given reason.
func (in *AddOnRevision) ClearAppliedImageConfigRef(reason pkgv1.ImageConfigRefReason) {
	for i := range in.Status.AppliedImageConfigRefs {
		if in.Status.AppliedImageConfigRefs[i].Reason == reason {
			in.Status.AppliedImageConfigRefs = append(in.Status.AppliedImageConfigRefs[:i], in.Status.AppliedImageConfigRefs[i+1:]...)
			return
		}
	}
}

// GetResolvedSource returns the resolved source package.
func (in *AddOnRevision) GetResolvedSource() string {
	return in.Status.ResolvedPackage
}

// SetResolvedSource sets the resolved source package.
func (in *AddOnRevision) SetResolvedSource(s string) {
	in.Status.ResolvedPackage = s
}

// SetConditions sets the status conditions for the AddOnRevision.
func (in *AddOnRevision) SetConditions(c ...xpv1.Condition) {
	in.Status.SetConditions(c...)
}

// GetCondition returns the condition for the given ConditionType if it exists,
// otherwise returns an empty condition.
func (in *AddOnRevision) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return in.Status.GetCondition(ct)
}

// CleanConditions removes all conditions from the AddOnRevision.
func (in *AddOnRevision) CleanConditions() {
	in.Status.Conditions = []xpv1.Condition{}
}

// GetObjects returns the objects associated with the AddOnRevision.
func (in *AddOnRevision) GetObjects() []xpv1.TypedReference {
	return in.Status.ObjectRefs
}

// SetObjects sets the objects associated with the AddOnRevision.
func (in *AddOnRevision) SetObjects(c []xpv1.TypedReference) {
	in.Status.ObjectRefs = c
}

// GetSource returns the package source.
func (in *AddOnRevision) GetSource() string {
	return in.Spec.Package
}

// SetSource sets the package source.
func (in *AddOnRevision) SetSource(s string) {
	in.Spec.Package = s
}

// GetPackagePullSecrets returns the package pull secrets.
func (in *AddOnRevision) GetPackagePullSecrets() []corev1.LocalObjectReference {
	return in.Spec.PackagePullSecrets
}

// SetPackagePullSecrets sets the package pull secrets.
func (in *AddOnRevision) SetPackagePullSecrets(s []corev1.LocalObjectReference) {
	in.Spec.PackagePullSecrets = s
}

// GetPackagePullPolicy returns the package pull policy.
func (in *AddOnRevision) GetPackagePullPolicy() *corev1.PullPolicy {
	return in.Spec.PackagePullPolicy
}

// SetPackagePullPolicy sets the package pull policy.
func (in *AddOnRevision) SetPackagePullPolicy(i *corev1.PullPolicy) {
	in.Spec.PackagePullPolicy = i
}

// GetDesiredState returns the desired state of the package revision.
func (in *AddOnRevision) GetDesiredState() pkgv1.PackageRevisionDesiredState {
	return in.Spec.DesiredState
}

// SetDesiredState sets the desired state of the package revision.
func (in *AddOnRevision) SetDesiredState(d pkgv1.PackageRevisionDesiredState) {
	in.Spec.DesiredState = d
}

// GetIgnoreCrossplaneConstraints returns whether to ignore crossplane constraints.
func (in *AddOnRevision) GetIgnoreCrossplaneConstraints() *bool {
	return in.Spec.IgnoreCrossplaneConstraints
}

// SetIgnoreCrossplaneConstraints sets whether to ignore crossplane constraints.
func (in *AddOnRevision) SetIgnoreCrossplaneConstraints(b *bool) {
	in.Spec.IgnoreCrossplaneConstraints = b
}

// GetRevision returns the revision number.
func (in *AddOnRevision) GetRevision() int64 {
	return in.Spec.Revision
}

// SetRevision sets the revision number.
func (in *AddOnRevision) SetRevision(r int64) {
	in.Spec.Revision = r
}

// GetSkipDependencyResolution returns whether to skip dependency resolution.
func (in *AddOnRevision) GetSkipDependencyResolution() *bool {
	return in.Spec.SkipDependencyResolution
}

// SetSkipDependencyResolution sets whether to skip dependency resolution.
func (in *AddOnRevision) SetSkipDependencyResolution(skip *bool) {
	in.Spec.SkipDependencyResolution = skip
}

// GetDependencyStatus returns the dependency status.
func (in *AddOnRevision) GetDependencyStatus() (found, installed, invalid int64) {
	return in.Status.FoundDependencies, in.Status.InstalledDependencies, in.Status.InvalidDependencies
}

// SetDependencyStatus sets the dependency status.
func (in *AddOnRevision) SetDependencyStatus(found, installed, invalid int64) {
	in.Status.FoundDependencies = found
	in.Status.InstalledDependencies = installed
	in.Status.InvalidDependencies = invalid
}

// GetCommonLabels returns the common labels.
func (in *AddOnRevision) GetCommonLabels() map[string]string {
	return in.Spec.CommonLabels
}

// SetCommonLabels sets the common labels.
func (in *AddOnRevision) SetCommonLabels(l map[string]string) {
	in.Spec.CommonLabels = l
}

// GetRuntimeConfigRef returns the runtime config reference.
func (in *AddOnRevision) GetRuntimeConfigRef() *pkgv1.RuntimeConfigReference {
	if in.Spec.RuntimeConfigReference == nil {
		return nil
	}
	return &pkgv1.RuntimeConfigReference{
		APIVersion: in.Spec.RuntimeConfigReference.APIVersion,
		Kind:       in.Spec.RuntimeConfigReference.Kind,
		Name:       in.Spec.RuntimeConfigReference.Name,
	}
}

// SetRuntimeConfigRef sets the runtime config reference.
func (in *AddOnRevision) SetRuntimeConfigRef(r *pkgv1.RuntimeConfigReference) {
	in.Spec.RuntimeConfigReference = &RuntimeConfigReference{
		APIVersion: r.APIVersion,
		Kind:       r.Kind,
		Name:       r.Name,
	}
}

// GetTLSServerSecretName returns the TLS server secret name.
func (in *AddOnRevision) GetTLSServerSecretName() *string {
	return nil
}

// SetTLSServerSecretName sets the TLS server secret name.
func (in *AddOnRevision) SetTLSServerSecretName(_ *string) {}

// GetTLSClientSecretName returns the TLS client secret name.
func (in *AddOnRevision) GetTLSClientSecretName() *string {
	return nil
}

// SetTLSClientSecretName sets the TLS client secret name.
func (in *AddOnRevision) SetTLSClientSecretName(_ *string) {}

// Implement XP Revision List interface for AddOnRevisionList.
var _ pkgv1.PackageRevisionList = &AddOnRevisionList{}

// GetRevisions returns the list of package revisions.
func (in *AddOnRevisionList) GetRevisions() []pkgv1.PackageRevision {
	prs := make([]pkgv1.PackageRevision, len(in.Items))
	for i, r := range in.Items {
		prs[i] = &r
	}
	return prs
}
