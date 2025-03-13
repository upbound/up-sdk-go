// Copyright 2025 Upbound Inc
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
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ObjectParameters are the configurable fields of a Object.
type ObjectParameters struct {
	// Raw JSON representation of the kubernetes object to be created.
	// +kubebuilder:validation:EmbeddedResource
	// +kubebuilder:pruning:PreserveUnknownFields
	Manifest runtime.RawExtension `json:"manifest"`
}

// ObjectObservation are the observable fields of a Object.
type ObjectObservation struct {
	// Raw JSON representation of the remote object.
	// +kubebuilder:validation:EmbeddedResource
	// +kubebuilder:pruning:PreserveUnknownFields
	Manifest runtime.RawExtension `json:"manifest,omitempty"`
}

// CompositeReference points to the composite that holds the reference.
type CompositeReference struct {
	// APIVersion of the referencing composite resource.
	// +required
	APIVersion string `json:"apiVersion"`
	// Kind of the referencing composite resource.
	// +required
	// +kubebuilder:validation:MinLength=1
	Kind string `json:"kind"`
	// Name of the referencing composite resource.
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
}

// CompositeReferencePath points to the composite that holds the reference
// at a specific JSONPath.
type CompositeReferencePath struct {
	CompositeReference `json:",inline"`

	// JSONPath is where the reference is stored in the composite.
	// The reference is resolved in the context of the claim.
	// The JSONPath is expected to end with a object field called `*Ref`, or
	// an array field called `*Refs`. In the former case, the object field must
	// have the following schema. In latter case, each item in the array must:
	//
	// apiVersion: string
	// kind: string
	// name: string
	// namespace: string # optional, defaults to be a local reference.
	// UID: string # optional
	// grants: []string # optional, defaults to ["Observe"], but can hold any of
	//                    ["Observe", "Create", "Update", "Delete"].
	// claims: []string # optional, additional verbs the user must be authorized to.
	//
	// Grants are a closed-set, corresponding to the management policy of
	// Crossplane. Grants are authorized through RBAC and Upbound ReBAC for the
	// referenced object when creating or updating a reference.
	//
	// Claims are optional verbs the user creating or referencing an object
	// must be authorized to. These can include custom verbs.
	//
	// With grants the user specifies the permissions that the claim gives to the
	// composite: "Observe" allows the composite to read the object, "Create" to
	// create it, "Update" to update it, and "Delete" to delete it.
	//
	// Claims tell the composition author about additional permissions the user
	// might have.
	//
	// OpenAPI v3 validation and defaulting can be used to restrict and/or
	// auto-populate the fields.
	//
	// A JSONPath must have a counterpart schema in the CompositeResourceDefinition's
	// references.upbound.io/schema annotation.
	//
	// +required
	// +kubebuilder:validation:MinLength=1
	JSONPath string `json:"jsonPath"`
}

// A ObjectSpec defines the desired state of a Object.
type ObjectSpec struct {
	// THIS IS A BETA FIELD. It is on by default but can be opted out
	// through a Crossplane feature flag.
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
	// when this managed resource is deleted - either "Delete" or "Orphan" the
	// external resource.
	// This field is planned to be deprecated in favor of the ManagementPolicies
	// field in a future release. Currently, both could be set independently and
	// non-default values would be honored if the feature flag is enabled.
	// See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223
	// +optional
	// +kubebuilder:default=Delete
	DeletionPolicy xpv1.DeletionPolicy `json:"deletionPolicy,omitempty"`

	// Readiness defines how the object's readiness condition should be computed.
	// +optional
	// +kubebuilder:default={}
	Readiness Readiness `json:"readiness,omitempty"`

	// Composite is the composite object that holds the reference.
	// The composite must be bound to a claim, that can be local or remote.
	// +required
	Composite CompositeReferencePath `json:"composite"`

	// OwnerPolicy defines how created objects should be owned.
	// 'OnCreate' requires management policy 'Create' or '*'.
	// +optional
	OwnerPolicy OwnerPolicy `json:"ownerPolicy,omitempty"`

	// ForProvider is the object's desired state.
	// +optional
	ForProvider ObjectParameters `json:"forProvider"`
}

// OwnerPolicy defines how the created objects' owner should be computed.
// +kubebuilder:validation:Enum=OnCreate
type OwnerPolicy string

const (
	// OwnerPolicyOnCreate means the created object should be owned by the claim.
	OwnerPolicyOnCreate OwnerPolicy = "OnCreate"
)

// ReadinessPolicy defines how the ReferencedObject's readiness condition should
// be computed.
type ReadinessPolicy string

const (
	// ReadinessPolicyWhenSynced means the ReferencedObject is marked as ready
	// when the referenced object has been synced. This includes the case when
	// the reference does not exist.
	ReadinessPolicyWhenSynced ReadinessPolicy = "WhenSynced"
	// ReadinessPolicyObjectExists means the ReferencedObject is marked as ready
	// when the referenced object exists.
	ReadinessPolicyObjectExists ReadinessPolicy = "ObjectExists"
	// ReadinessPolicyDeriveFromObject means the ReferencedObject is marked as
	// ready if and only if the referenced object is considered ready.
	ReadinessPolicyDeriveFromObject ReadinessPolicy = "ObjectReady"
	// ReadinessPolicyAllTrue means that all conditions have status true on the
	// referenced object. There must be at least one condition.
	ReadinessPolicyAllTrue ReadinessPolicy = "ObjectConditionsAllTrue"

	// TODO(sttts): have more policies or extra knobs to be ready when no reference exists, or the object does not exist.
)

// Readiness defines how the object's readiness condition should be computed,
// if not specified it will be considered ready as soon as the underlying external
// resource is considered up-to-date.
type Readiness struct {
	// Policy defines how the Object's readiness condition should be computed.
	// +optional
	// +kubebuilder:validation:Enum=WhenSynced;ObjectExists;ObjectReady;ObjectConditionsAllTrue
	// +kubebuilder:default=ObjectExists
	Policy ReadinessPolicy `json:"policy,omitempty"`
}

// A ObjectStatus represents the observed state of a Object.
type ObjectStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// AtProvider is the object's observed state.
	// +optional
	AtProvider ObjectObservation `json:"atProvider,omitempty"`
}

// A ReferencedObject represents a Kubernetes object that is referenced by a
// claim.
//
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="KIND",type="string",JSONPath=".spec.forProvider.manifest.kind"
// +kubebuilder:printcolumn:name="KIND",type="string",JSONPath=".spec.composite.kind",priority=1
// +kubebuilder:printcolumn:name="APIVERSION",type="string",JSONPath=".spec.composite.apiVersion",priority=2
// +kubebuilder:printcolumn:name="COMPOSITE",type="string",JSONPath=".spec.composite.name",priority=1
// +kubebuilder:printcolumn:name="JSONPATH",type="string",JSONPath=".spec.composite.JSONPath",priority=1
// +kubebuilder:printcolumn:name="REFERENCEKIND",type="string",JSONPath=".status.atProvider.manifest.kind",priority=1
// +kubebuilder:printcolumn:name="REFERENCEAPIVERSION",type="string",JSONPath=".status.atProvider.manifest.apiVersion",priority=2
// +kubebuilder:printcolumn:name="REFERENCEDNAME",type="string",JSONPath=".status.atProvider.manifest.metadata.name",priority=1
// +kubebuilder:printcolumn:name="REFERENCEDNAMESPACE",type="string",JSONPath=".status.atProvider.manifest.metadata.namespace",priority=1
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={managed,kubernetes}
// +kubebuilder:storageversion
type ReferencedObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObjectSpec   `json:"spec"`
	Status ObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReferencedObjectList contains a list of ReferencedObjects.
type ReferencedObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReferencedObject `json:"items"`
}

// GroupVersionKind returns the GroupVersionKind of the ReferencedObject's
// composite reference.
func (r *CompositeReference) GroupVersionKind() schema.GroupVersionKind {
	cs := strings.Split(r.APIVersion, "/")
	if len(cs) == 1 {
		return schema.GroupVersionKind{
			Group:   "",
			Version: r.APIVersion,
			Kind:    r.Kind,
		}
	}
	return schema.GroupVersionKind{
		Group:   cs[0],
		Version: cs[1],
		Kind:    r.Kind,
	}
}

// GetCondition of this ReferencedObject.
func (mg *ReferencedObject) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ReferencedObject.
func (mg *ReferencedObject) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ReferencedObject.
func (mg *ReferencedObject) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// SetConditions of this ReferencedObject.
func (mg *ReferencedObject) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ReferencedObject.
func (mg *ReferencedObject) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ReferencedObject.
func (mg *ReferencedObject) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// ReferencedObjectKind is the kind of the ReferencedObject.
var ReferencedObjectKind = reflect.TypeOf(ReferencedObject{}).Name()

// ReferencedObjectListKind is the kind of a list of ReferencedObjects.
var ReferencedObjectListKind = reflect.TypeOf(ReferencedObjectList{}).Name()

func init() {
	SchemeBuilder.Register(&ReferencedObject{}, &ReferencedObjectList{})
}
