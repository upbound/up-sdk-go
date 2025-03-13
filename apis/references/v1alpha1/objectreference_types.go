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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ObjectReference holds the reference to another object.
type ObjectReference struct {
	xpv1.TypedReference `json:",inline"`

	// Namespace is the namespace of the referenced object, if the given kind
	// is namespaced.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Grants are the permissions on the referenced object granted to the
	// composition author. The default is ["get"|. It is a closed set with the
	// possible values "Observe", "Create", "Update", "Delete", "*".
	// The grants are authorized against the referenced object when the
	// reference is created or updated by the user.
	// +optional
	Grants []xpv1.ManagementAction `json:"grants,omitempty"`
}

// ObjectsReference holds the reference to multiple objects.
type ObjectsReference struct {
	// APIVersion of the referenced objects.
	APIVersion string `json:"apiVersion"`

	// Kind of the referenced objects.
	Kind string `json:"kind"`

	// Namespace is the namespace of the referenced objects, if the given kind
	// is namespaced.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// MatchLabels selects the objects that should be referenced.
	MatchLabels metav1.LabelSelector `json:"matchLabels,omitempty"`

	// JSONPath is an optional RFC 9535 JSONPath expression that specifies which
	// part of the referenced objects should be referenced.
	JSONPath string `json:"jsonPath,omitempty"`
}
