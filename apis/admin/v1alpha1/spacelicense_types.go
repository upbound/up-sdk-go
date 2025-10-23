/*
Copyright 2025 The Upbound Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

const (
	// SpaceLicenseName is the required name for the singleton SpaceLicense resource.
	SpaceLicenseName = "space"

	// SpaceLicenseSecretName is the name of the secret that contains the license file.
	SpaceLicenseSecretName = "space-license"

	// SpaceLicenseSecretKeyDefault is the default key within the secret that contains the license file.
	SpaceLicenseSecretKeyDefault = "license.json"

	// PlanCommunity is the plan name for community edition licenses.
	PlanCommunity = "community"
)

// SpaceLicenseSecretRef references a Kubernetes secret containing the license file.
type SpaceLicenseSecretRef struct {
	// Name of the secret containing the license.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Namespace of the secret containing the license.
	// +kubebuilder:validation:Required
	Namespace string `json:"namespace"`

	// Key within the secret that contains the license file.
	// Defaults to "license.json".
	// +kubebuilder:default="license.json"
	// +optional
	Key string `json:"key,omitempty"`
}

// SpaceLicenseSpec defines the desired state of SpaceLicense.
type SpaceLicenseSpec struct {
	// SecretRef references the secret that contains the license file.
	// If not provided, the license will be considered as community edition.
	// +optional
	SecretRef *SpaceLicenseSecretRef `json:"secretRef,omitempty"`
}

// SpaceLicenseCapacity defines the capacity granted by the license.
type SpaceLicenseCapacity struct {
	// ResourceHours is the total number of resource hours granted by the license.
	// +optional
	ResourceHours int64 `json:"resourceHours,omitempty"`

	// Operations is the total number of operations granted by the license.
	// +optional
	Operations int64 `json:"operations,omitempty"`
}

// SpaceLicenseRestrictions defines environmental or operational constraints for the license.
type SpaceLicenseRestrictions struct {
	// ClusterUUID specifies the cluster UUID that this license is restricted to.
	// If present, the license is only valid for the cluster with this UUID.
	// +optional
	ClusterUUID string `json:"clusterUUID,omitempty"`
}

// CrossplaneStatus defines the status of the current Crossplane installation.
type CrossplaneStatus struct {
	// Version is the version of Crossplane.
	// +optional
	Version string `json:"version,omitempty"`
}

// InstalledComponentsStatus defines the status of the installed components.
type InstalledComponentsStatus struct {
	// Crossplane contains the status of the current Crossplane installation.
	// +optional
	Crossplane *CrossplaneStatus `json:"crossplane,omitempty"`
}

// SpaceLicenseUsage defines the cumulative usage data tracked by the system.
type SpaceLicenseUsage struct {
	// ResourceHours is the total number of resource hours consumed.
	// +optional
	ResourceHours int64 `json:"resourceHours"`

	// Operations is the total number of operations consumed.
	// +optional
	Operations int64 `json:"operations"`

	// ResourceHoursUtilization is the percentage of resource hours capacity used.
	// +optional
	ResourceHoursUtilization string `json:"resourceHoursUtilization"`

	// OperationsUtilization is the percentage of operations capacity used.
	// +optional
	OperationsUtilization string `json:"operationsUtilization"`

	// FirstMeasurement is the timestamp of the first usage measurement.
	// +optional
	FirstMeasurement *metav1.Time `json:"firstMeasurement,omitempty"`

	// LastMeasurement is the timestamp of the last usage measurement.
	// +optional
	LastMeasurement *metav1.Time `json:"lastMeasurement,omitempty"`
}

// SpaceLicenseStatus defines the observed state of SpaceLicense.
type SpaceLicenseStatus struct {
	xpv1.ConditionedStatus `json:",inline"`

	// ID is the ID associated with the license.
	// This field is owned by the license-check-controller.
	// +optional
	ID string `json:"id,omitempty"`

	// Plan is the commercial plan associated with the license (e.g., "community", "standard").
	// This field is owned by the license-check-controller.
	// +optional
	Plan string `json:"plan,omitempty"`

	// Capacity contains the capacity granted by the license.
	// This entire block is owned by the license-check-controller.
	// +optional
	Capacity *SpaceLicenseCapacity `json:"capacity,omitempty"`

	// EnabledFeatures lists the commercial features enabled by this license.
	// This field is owned by the license-check-controller.
	// +optional
	EnabledFeatures []string `json:"enabledFeatures,omitempty"`

	// CreatedAt is when the license was created.
	// This field is owned by the license-check-controller.
	// +optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	// ExpiresAt is when the license expires.
	// This field is owned by the license-check-controller.
	// +optional
	ExpiresAt *metav1.Time `json:"expiresAt,omitempty"`

	// GracePeriodEndsAt is when the license grace period ends and the license becomes completely invalid.
	// This field is owned by the license-check-controller.
	// +optional
	GracePeriodEndsAt *metav1.Time `json:"gracePeriodEndsAt,omitempty"`

	// Restrictions contains any usage restrictions associated with the license.
	// This field is owned by the license-check-controller.
	// +optional
	Restrictions *SpaceLicenseRestrictions `json:"restrictions,omitempty"`

	// Usage contains the cumulative usage data tracked by the system.
	// This entire block is owned by the metering-controller.
	// +optional
	Usage *SpaceLicenseUsage `json:"usage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories=spaces
// +kubebuilder:printcolumn:name="PLAN",type="string",JSONPath=".status.plan"
// +kubebuilder:printcolumn:name="VALID",type="string",JSONPath=".status.conditions[?(@.type=='LicenseValid')].status"
// +kubebuilder:printcolumn:name="REASON",type="string",JSONPath=".status.conditions[?(@.type=='LicenseValid')].reason"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"

// SpaceLicense is the resource for Space license configuration.
type SpaceLicense struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpaceLicenseSpec   `json:"spec,omitempty"`
	Status SpaceLicenseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpaceLicenseList contains a list of SpaceLicense.
type SpaceLicenseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpaceLicense `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpaceLicense{}, &SpaceLicenseList{})
}
