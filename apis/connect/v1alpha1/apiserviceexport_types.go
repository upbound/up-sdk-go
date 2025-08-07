// Copyright 2024 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	"reflect"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

// APIServiceExport specifies the resource to be exported. Its spec
// is that of a CRD without webhooks.
// This resource lives in a service provider group of a Space.
// The CRD can be sourced from a control plane in the same Spaces group that
// the export lives.
//
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Description",type="string",JSONPath=".spec.description"
// +kubebuilder:printcolumn:name="Group",type="string",JSONPath=".status.apiGroup"
// +kubebuilder:printcolumn:name="Resource",type="string",JSONPath=".status.names.plural"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=`.status.message`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories={connect},shortName={export,exports}
type APIServiceExport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +required
	Spec   APIServiceExportSpec   `json:"spec,omitempty"`
	Status APIServiceExportStatus `json:"status,omitempty"`
}

// APIServiceExportList contains a list of APIServiceExports.
//
// +kubebuilder:object:root=true
type APIServiceExportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIServiceExport `json:"items"`
}

// APIServiceExportSpec defines the desired state of a APIServiceExport.
type APIServiceExportSpec struct {
	// apiGroup is the group name for the exported API
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern="^[a-z0-9]+(.[a-z0-9-]+)*$"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="apiGroup is immutable"
	APIGroup string `json:"apiGroup"`

	// resource is the name of the exported API resource
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="resource is immutable"
	Resource string `json:"resource"`

	// description is a human-readable description of the exported API.
	// It should be a short one-liner, with a maximum length of 48 chars.
	// +optional
	// +kubebuilder:validation:items:MaxLength=48
	Description string `json:"description,omitempty"`
	// apiSource is a reference to control plane that is the
	// source of truth of the CRD schema.
	// It can only be a control plane in the same Spaces group
	//
	// +required
	APISource APIServiceExportSource `json:"apiSource"`

	// enableLocatorsAtClaims controls whether to set locator annotations
	// on claims, that provides information about the control plane
	// that manages the claim.
	//
	// When true, sets the following annotations on the claim:
	//   connect.upbound.io/namespace
	//   connect.upbound.io/controlplane
	//   connect.upbound.io/group
	//   connect.upbound.io/space
	//
	// +kubebuilder:default=false
	EnableLocatorsAtClaims bool `json:"enableLocatorsAtClaims"`
}

// APIServiceExportCRDSpec is an apiextensionsv1.CustomResourceDefinitionSpec
// equivalent, that represents the CRD to export.
type APIServiceExportCRDSpec struct {
	// group is the API group of the defined custom resource.
	// k8s core API is not supported.
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern="^[a-z0-9]+(.[a-z0-9-]+)*$"
	APIGroup string `json:"apiGroup"`

	// names specify the resource and kind names for the custom resource.
	//
	// +required
	Names apiextensionsv1.CustomResourceDefinitionNames `json:"names"`

	// scope indicates whether the defined custom resource is cluster- or namespace-scoped.
	// Allowed values are `Cluster` and `Namespaced`.
	//
	// +optional
	// +kubebuilder:validation:Enum=Cluster;Namespaced
	// +kubebuilder:default=Namespaced
	Scope apiextensionsv1.ResourceScope `json:"scope"`

	// versions is the API version of the defined custom resource.
	//
	// +required
	// +listType=map
	// +listMapKey=name
	// +kubebuilder:validation:MinItems=1
	Versions []APIServiceExportVersion `json:"versions"`
}

// APIServiceExportVersion describes one API version of a resource.
type APIServiceExportVersion struct {
	// name is the version name, e.g. “v1”, “v2beta1”, etc.
	// The custom resources are served under this version at `/apis/<group>/<version>/...` if `served` is true.
	//
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=^v[1-9][0-9]*([a-z]+[1-9][0-9]*)?$
	Name string `json:"name"`
	// storage indicates this version should be used when persisting custom resources to storage.
	// There must be exactly one version with storage=true.
	//
	// +required
	// +kubebuilder:validation:Required
	Storage bool `json:"storage"`
	//nolint:gocritic // linter thinks below is a deprecation comment, it is a field
	// deprecated indicates this version of the custom resource API is deprecated.
	// When set to true, API requests to this version receive a warning header in the server response.
	// Defaults to false.
	//
	// +optional
	Deprecated bool `json:"deprecated,omitempty"`
	// deprecationWarning overrides the default warning returned to API clients.
	// May only be set when `deprecated` is true.
	// The default warning indicates this version is deprecated and recommends use
	// of the newest served version of equal or greater stability, if one exists.
	//
	// +optional
	DeprecationWarning *string `json:"deprecationWarning,omitempty"`
	// schema describes the structural schema used for validation, pruning, and defaulting
	// of this version of the custom resource.
	//
	// +required
	// +kubebuilder:validation:Required
	Schema APIServiceExportSchema `json:"schema"`
	// subresources specify what subresources this version of the defined custom resource have.
	//
	// +optional
	Subresources apiextensionsv1.CustomResourceSubresources `json:"subresources,omitempty"`
	// additionalPrinterColumns specifies additional columns returned in Table output.
	// See https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables for details.
	// If no columns are specified, a single column displaying the age of the custom resource is used.
	//
	// +optional
	// +listType=map
	// +listMapKey=name
	AdditionalPrinterColumns []apiextensionsv1.CustomResourceColumnDefinition `json:"additionalPrinterColumns,omitempty"`
}

// APIServiceExportSchema describes the validation schema of an API
type APIServiceExportSchema struct {
	// openAPIV3Schema is the OpenAPI v3 schema to use for validation and pruning.
	//
	// +kubebuilder:pruning:PreserveUnknownFields
	// +structType=atomic
	// +required
	// +kubebuilder:validation:Required
	OpenAPIV3Schema runtime.RawExtension `json:"openAPIV3Schema"`
}

// APIServiceExportSource defines the source of the exported API.
// It must refer to a control plane in the same Spaces group
// with the APIServiceExport.
type APIServiceExportSource struct {
	// ControlPlane is a reference to a control plane in the same Spaces group
	// +required
	ControlPlane ControlPlaneNameRef `json:"controlPlane"`
}

// ControlPlaneNameRef identifies a control plane with its name.
type ControlPlaneNameRef struct {
	// Name is the name of the control plane
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
}

// APIServiceExportStatus defines the observed state of a APIServiceExport.
type APIServiceExportStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	APIServiceExportCRDSpec `json:",inline"`
	// Message provides human-readable information about the current status of
	// the APIServiceExport.
	//
	// +optional
	Message string `json:"message,omitempty"`
}

var (
	// APIServiceExportKind is kind of APIServiceExport
	APIServiceExportKind = reflect.TypeOf(APIServiceExport{}).Name()
)

func init() {
	SchemeBuilder.Register(&APIServiceExport{}, &APIServiceExportList{})
}
