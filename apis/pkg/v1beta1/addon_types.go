// Copyright 2025 Upbound Inc
// All rights reserved

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	pkgv1 "github.com/crossplane/crossplane/apis/pkg/v1"
)

// +kubebuilder:object:root=true
// +genclient
// +genclient:nonNamespaced

// An AddOn installs an OCI compatible Upbound package, extending a Control
// Plane with new capabilities.
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="INSTALLED",type="string",JSONPath=".status.conditions[?(@.type=='Installed')].status"
// +kubebuilder:printcolumn:name="HEALTHY",type="string",JSONPath=".status.conditions[?(@.type=='Healthy')].status"
// +kubebuilder:printcolumn:name="PACKAGE",type="string",JSONPath=".spec.package"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={upbound,pkg}
type AddOn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AddOnSpec   `json:"spec"`
	Status AddOnStatus `json:"status,omitempty"`
}

// AddOnSpec specifies the configuration of an AddOn.
type AddOnSpec struct {
	pkgv1.PackageSpec  `json:",inline"`
	PackageRuntimeSpec `json:",inline"`
}

// AddOnStatus represents the observed state of an AddOn.
type AddOnStatus struct {
	xpv1.ConditionedStatus `json:",inline"`
	pkgv1.PackageStatus    `json:",inline"`
}

// +kubebuilder:object:root=true

// AddOnList contains a list of AddOn.
type AddOnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddOn `json:"items"`
}

// PackageRuntimeSpec specifies configuration for the runtime of a package.
type PackageRuntimeSpec struct {
	// RuntimeConfigRef references a RuntimeConfig resource that will be used
	// to configure the package runtime.
	// +optional
	// +kubebuilder:default={"name": "default"}
	RuntimeConfigReference *RuntimeConfigReference `json:"runtimeConfigRef,omitempty"`
}

// A RuntimeConfigReference to a runtime config resource that will be used
// to configure the package runtime.
type RuntimeConfigReference struct {
	// API version of the referent.
	// +optional
	// +kubebuilder:default="pkg.upbound.io/v1alpha1"
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind of the referent.
	// +optional
	// +kubebuilder:default="AddOnRuntimeConfig"
	Kind *string `json:"kind,omitempty"`
	// Name of the RuntimeConfig.
	Name string `json:"name"`
}

// Implement XP Package interface for AddOn.
var _ pkgv1.Package = &AddOn{}

// GetAppliedImageConfigRefs returns the applied image config references.
func (in *AddOn) GetAppliedImageConfigRefs() []pkgv1.ImageConfigRef {
	return in.Status.AppliedImageConfigRefs
}

// SetAppliedImageConfigRefs sets the applied image config references.
func (in *AddOn) SetAppliedImageConfigRefs(refs ...pkgv1.ImageConfigRef) {
	in.Status.AppliedImageConfigRefs = refs
}

// ClearAppliedImageConfigRef clears the applied image config reference for a given reason.
func (in *AddOn) ClearAppliedImageConfigRef(reason pkgv1.ImageConfigRefReason) {
	for i := range in.Status.AppliedImageConfigRefs {
		if in.Status.AppliedImageConfigRefs[i].Reason == reason {
			in.Status.AppliedImageConfigRefs = append(in.Status.AppliedImageConfigRefs[:i], in.Status.AppliedImageConfigRefs[i+1:]...)
			return
		}
	}
}

// GetResolvedSource returns the resolved source package.
func (in *AddOn) GetResolvedSource() string {
	return in.Status.ResolvedPackage
}

// SetResolvedSource sets the resolved source package.
func (in *AddOn) SetResolvedSource(s string) {
	in.Status.ResolvedPackage = s
}

// SetConditions sets the status conditions for the AddOn.
func (in *AddOn) SetConditions(c ...xpv1.Condition) {
	in.Status.SetConditions(c...)
}

// GetCondition returns the condition for the given ConditionType if it exists,
// otherwise returns an empty condition.
func (in *AddOn) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return in.Status.GetCondition(ct)
}

// CleanConditions removes all conditions from the AddOn.
func (in *AddOn) CleanConditions() {
	in.Status.Conditions = []xpv1.Condition{}
}

// GetSource returns the package source.
func (in *AddOn) GetSource() string {
	return in.Spec.Package
}

// SetSource sets the package source.
func (in *AddOn) SetSource(s string) {
	in.Spec.Package = s
}

// GetActivationPolicy returns the revision activation policy.
func (in *AddOn) GetActivationPolicy() *pkgv1.RevisionActivationPolicy {
	return in.Spec.RevisionActivationPolicy
}

// SetActivationPolicy sets the revision activation policy.
func (in *AddOn) SetActivationPolicy(a *pkgv1.RevisionActivationPolicy) {
	in.Spec.RevisionActivationPolicy = a
}

// GetPackagePullSecrets returns the package pull secrets.
func (in *AddOn) GetPackagePullSecrets() []corev1.LocalObjectReference {
	return in.Spec.PackagePullSecrets
}

// SetPackagePullSecrets sets the package pull secrets.
func (in *AddOn) SetPackagePullSecrets(s []corev1.LocalObjectReference) {
	in.Spec.PackagePullSecrets = s
}

// GetPackagePullPolicy returns the package pull policy.
func (in *AddOn) GetPackagePullPolicy() *corev1.PullPolicy {
	return in.Spec.PackagePullPolicy
}

// SetPackagePullPolicy sets the package pull policy.
func (in *AddOn) SetPackagePullPolicy(i *corev1.PullPolicy) {
	in.Spec.PackagePullPolicy = i
}

// GetRevisionHistoryLimit returns the revision history limit.
func (in *AddOn) GetRevisionHistoryLimit() *int64 {
	return in.Spec.RevisionHistoryLimit
}

// SetRevisionHistoryLimit sets the revision history limit.
func (in *AddOn) SetRevisionHistoryLimit(l *int64) {
	in.Spec.RevisionHistoryLimit = l
}

// GetIgnoreCrossplaneConstraints returns whether to ignore crossplane constraints.
func (in *AddOn) GetIgnoreCrossplaneConstraints() *bool {
	return in.Spec.IgnoreCrossplaneConstraints
}

// SetIgnoreCrossplaneConstraints sets whether to ignore crossplane constraints.
func (in *AddOn) SetIgnoreCrossplaneConstraints(b *bool) {
	in.Spec.IgnoreCrossplaneConstraints = b
}

// GetCurrentRevision returns the current revision.
func (in *AddOn) GetCurrentRevision() string {
	return in.Status.CurrentRevision
}

// SetCurrentRevision sets the current revision.
func (in *AddOn) SetCurrentRevision(r string) {
	in.Status.CurrentRevision = r
}

// GetCurrentIdentifier returns the current identifier.
func (in *AddOn) GetCurrentIdentifier() string {
	return in.Status.CurrentIdentifier
}

// SetCurrentIdentifier sets the current identifier.
func (in *AddOn) SetCurrentIdentifier(r string) {
	in.Status.CurrentIdentifier = r
}

// GetSkipDependencyResolution returns whether to skip dependency resolution.
func (in *AddOn) GetSkipDependencyResolution() *bool {
	return in.Spec.SkipDependencyResolution
}

// SetSkipDependencyResolution sets whether to skip dependency resolution.
func (in *AddOn) SetSkipDependencyResolution(skip *bool) {
	in.Spec.SkipDependencyResolution = skip
}

// GetCommonLabels returns the common labels.
func (in *AddOn) GetCommonLabels() map[string]string {
	return in.Spec.CommonLabels
}

// SetCommonLabels sets the common labels.
func (in *AddOn) SetCommonLabels(l map[string]string) {
	in.Spec.CommonLabels = l
}

// GetRuntimeConfigRef returns the runtime config reference.
func (in *AddOn) GetRuntimeConfigRef() *pkgv1.RuntimeConfigReference {
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
func (in *AddOn) SetRuntimeConfigRef(r *pkgv1.RuntimeConfigReference) {
	in.Spec.RuntimeConfigReference = &RuntimeConfigReference{
		APIVersion: r.APIVersion,
		Kind:       r.Kind,
		Name:       r.Name,
	}
}

// GetTLSServerSecretName returns the TLS server secret name.
func (in *AddOn) GetTLSServerSecretName() *string {
	return nil
}

// GetTLSClientSecretName returns the TLS client secret name.
func (in *AddOn) GetTLSClientSecretName() *string {
	return nil
}
