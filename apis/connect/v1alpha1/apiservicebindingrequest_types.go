// Copyright 2024 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	"reflect"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// APIServiceBindingRequest represents the binding requests of a
// consumer cluster and provides the bindable exports in a Spaces group
// and scheduling information for each consumer cluster namespace.
// This object lives in the service provider Upbound Space cluster, inside a group,
// alongside the APIServiceExports. Created per consumer cluster in every group
// it consumes API services from, named the same as the ClusterBinding.
//
// This object lives in the service provider Upbound Space.
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Exports",type="string",JSONPath=`.status.message`
// +kubebuilder:printcolumn:name="Namespaces",type="string",JSONPath=`.status.message`
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=`.status.message`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories={connect},shortName={bindingrequest,bindingrequests}
type APIServiceBindingRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +required
	Spec   APIServiceBindingRequestSpec   `json:"spec,omitempty"`
	Status APIServiceBindingRequestStatus `json:"status,omitempty"`
}

// APIServiceBindingRequestList contains a list of APIServiceBindingRequest.
// +kubebuilder:object:root=true
type APIServiceBindingRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIServiceBindingRequest `json:"items"`
}

// APIServiceBindingRequestSpec defines the desired state of APIServiceBindingRequest.
type APIServiceBindingRequestSpec struct {
	// bindableExports is the list of bindable APIServiceExports for the consumer
	// cluster. A user adding a bindable export needs verb=bind
	// permissions against the APIServiceExport.
	// +listType=map
	// +listMapKey=name
	BindableExports []APIServiceExportRef `json:"bindableExports"`
	// namespaces defines a bound API consumption
	// +listType=map
	// +listMapKey=name
	Namespaces []NamespaceRequest `json:"namespaces"`
}

// APIServiceExportRef identifies an APIServiceExport
type APIServiceExportRef struct {
	// name is the name of the APIServiceExport
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
	// versions is the versions of the exported API
	// Only one version is allowed for now. More might be allowed in the future.
	// +required
	// +listType=set
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=1
	Versions []string `json:"versions"`
}

// NamespaceRequest represents a binding request of a consumer
// cluster namespace, and provides scheduling information
// for each resource
type NamespaceRequest struct {
	// name is the name of the namespace in the consumer cluster
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
	// resources is a list of export name and target namespace pairs,
	// that represents the exported API consumption in the consumer cluster
	// namespace and their corresponding target controlplane and namespace.
	// targets are owned and set by the scheduler according to the
	// scheduling logic.
	// +listType=map
	// +listMapKey=exportName
	// +kubebuilder:validation:items:MinLength=1
	Resources []ResourceTargetConfig `json:"resources"`
}

// ResourceTargetConfig represents a mapping between exports and
// target service implementor controlplane namespaces
type ResourceTargetConfig struct {
	// exportName is a reference to an APIServiceExport to configure a
	// target control plane for.
	// +required
	// +kubebuilder:validation:MinLength=1
	ExportName string `json:"exportName"`
	// target is the location that the claim will be scheduled to.
	// Points to a namespace of a control plane that can satisfy the claim.
	// This is owned by the scheduler at the service provider cluster
	// immutable after it is set
	//
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="target is immutable"
	Target *ResourceTarget `json:"target,omitempty"`
}

// ResourceTarget defines a target control plane
type ResourceTarget struct {
	// controlPlane defines the target control plane and a namespace
	// +required
	ControlPlane ControlPlaneNamespaceTarget `json:"controlPlane"`
}

// ControlPlaneNamespaceTarget is a reference to a namespace in a control plane
type ControlPlaneNamespaceTarget struct {
	// group is the Spaces group that the control plane belongs to
	// +required
	// +kubebuilder:validation:MinLength=1
	Group string `json:"group"`
	// name is the name of the control plane
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
	// namespace is the namespace name in the control plane
	// +required
	// +kubebuilder:validation:MinLength=1
	Namespace string `json:"namespace"`
}

// APIServiceBindingRequestStatus defines the observed state of a APIServiceBindingRequest.
// It reflects the status of the CRD of the consumer cluster.
type APIServiceBindingRequestStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// +listType=map
	// +listMapKey=name
	BoundExports []ExportStatus `json:"boundExports,omitempty"`

	// Message provides human-readable information about the current status of
	// the APIServiceBindingRequest.
	// +optional
	Message string `json:"message,omitempty"`
}

// ExportStatus defines the status of an exported API and corresponding CRD
type ExportStatus struct {
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
	// +required
	// +kubebuilder:validation:MinLength=1
	Resource string `json:"resource"`
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern="^[a-z0-9]+(.[a-z0-9-]+)*$"
	APIGroup string `json:"apiGroup"`
	// conditions indicate state for particular aspects of a CustomResourceDefinition
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []apiextensionsv1.CustomResourceDefinitionCondition `json:"conditions" protobuf:"bytes,1,opt,name=conditions"`
	// acceptedNames are the names that are actually being used to serve discovery.
	// They may be different from the names in spec.
	// +optional
	AcceptedNames apiextensionsv1.CustomResourceDefinitionNames `json:"acceptedNames"`
	// storedVersions lists all versions of CustomResources that were ever persisted. Tracking these
	// versions allows a migration path for stored versions in etcd. The field is mutable
	// so a migration controller can finish a migration to another version (ensuring
	// no old objects are left in storage), and then remove the rest of the
	// versions from this list.
	// Versions may not be removed from `spec.versions` while they exist in this list.
	// +optional
	BoundVersions []string `json:"boundVersions"`
	// versions are the exported versions
	// +listType=set
	Versions []string `json:"versions"`
}

var (
	// APIServiceBindingRequestKind is kind of APIServiceBindingRequest
	APIServiceBindingRequestKind = reflect.TypeOf(APIServiceBindingRequest{}).Name()
)

func init() {
	SchemeBuilder.Register(&APIServiceBindingRequest{}, &APIServiceBindingRequestList{})
}
