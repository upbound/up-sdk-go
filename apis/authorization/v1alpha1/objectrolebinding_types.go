// Copyright 2024 Upbound Inc
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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ObjectRoleBindingSpec is ObjectRoleBinding's spec.
type ObjectRoleBindingSpec struct {
	// Object references the object to which the listed subjects should have access at varying levels.
	// The object value is immutable after creation.
	// +kubebuilder:validation:Required
	Object Object `json:"object"`

	// Subjects should be a map type with both kind+name as a key
	// +listType=map
	// +listMapKey=kind
	// +listMapKey=name
	Subjects []SubjectBinding `json:"subjects"`
}

// Object represents the API object for which permissions are managed.
// In a ObjectRoleBinding context, the object exists in the same namespace
// as the referring ObjectRoleBinding, or the object is the namespace that
// the ObjectRoleBinding is in.
// In a SpaceObjectRoleBinding context, the object being pointed to must be
// non-namespaced.
type Object struct {
	// APIGroup defines the apiGroup of the object being pointed to.
	// With some minor differences, this is essentially matched as a DNS subdomain, like how Kubernetes validates it.
	// The Kubernetes legacy core group is denoted as "core".
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Pattern="^[a-z][a-z0-9-]{0,61}[a-z0-9](\\.[a-z][a-z0-9-]{0,61}[a-z0-9])*$"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="apiGroup is immutable"
	// +kubebuilder:validation:XValidation:rule="self == 'core'",message="apiGroup must be 'core' for now. This will change in the future."
	APIGroup string `json:"apiGroup"`

	// Resource defines the resource type (often kind in plural, e.g.
	// controlplanes) being pointed to.
	// With some minor differences, this is essentially matched as a DNS label, like how Kubernetes validates it.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern="^[a-z][a-z0-9-]{1,61}[a-z0-9]$"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="resource is immutable"
	// +kubebuilder:validation:XValidation:rule="self == 'namespaces'",message="resource must be 'namespaces' for now. This will change in the future."
	Resource string `json:"resource"`

	// Name points to the .metadata.name of the object targeted.
	// Kubernetes validates this as a DNS 1123 subdomain.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern="^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="name is immutable"
	Name string `json:"name"`
}

// SubjectKind Kind of subject being referenced.
type SubjectKind string

const (
	// SubjectKindUpboundTeam refers to an Upbound team as the subject kind.
	// For this kind, the name/identifier is the team UUID.
	SubjectKindUpboundTeam SubjectKind = "UpboundTeam"
)

// SubjectBinding contains a reference to the object or user identities a role
// binding applies to.
type SubjectBinding struct {
	// Kind of subject being referenced. Values defined by this API group are
	// for now only "UpboundTeam".
	// +kubebuilder:validation:Enum=UpboundTeam
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="self == 'UpboundTeam'",message="kind must be 'UpboundTeam' for now. This will change in the future."
	Kind SubjectKind `json:"kind"`

	// Name (identifier) of the subject (of the specified kind) being referenced.
	// The identifier must be 2-100 chars, [a-zA-Z0-9-], no repeating dashes, can't start/end with a dash.
	// Notably, a UUID fits that format.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=100
	// +kubebuilder:validation:Pattern="^([a-zA-Z0-9]+-?)+[a-zA-Z0-9]$"
	Name string `json:"name"`

	// Role this subject has on the associated Object.
	// The list of valid roles is defined for each target API resource separately.
	// For namespaces, valid values are "viewer", "editor", and "admin".
	// The format of this is essentially a RFC 1035 label with underscores instead of dashes, minimum three characters long.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern="^[a-z][a-z0-9_]{1,62}[a-z0-9]$"
	Role string `json:"role"`
}

// ObjectRoleBindingStatus is RoleBindings' status.
type ObjectRoleBindingStatus struct{}

// +kubebuilder:object:root=true

// A ObjectRoleBinding binds a namespaced API object to a set of subjects, at varying access levels.
// For now, there can be at most one ObjectRoleBinding pointing to each API object.
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories=iam
type ObjectRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   ObjectRoleBindingSpec   `json:"spec"`
	Status ObjectRoleBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectRoleBindingList contains a list of ObjectRoleBindings.
type ObjectRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectRoleBinding `json:"items"`
}

// ObjectRoleBindingKind is the kind of a ObjectRoleBinding.
//
//nolint:gochecknoglobals // This is an established pattern
var ObjectRoleBindingKind = reflect.TypeOf(ObjectRoleBinding{}).Name()

func init() {
	SchemeBuilder.Register(&ObjectRoleBinding{}, &ObjectRoleBindingList{})
}
