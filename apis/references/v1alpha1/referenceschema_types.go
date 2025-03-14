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
)

// Example:
// {
//   "apiVersion": "references.upbound.io/v1alpha1",
//   "kind": "ReferenceSchema",
//   "references": [{
//     "jsonPath": ".spec.secretRef",
//     "kinds": [{
//       "apiVersion": "v1",
//       "kind": "Secret"
//     }]
//   }]
// }

// ReferenceSchema hold the schema additions of references embedded in the
// XRD schema.
//
// +kubebuilder:object:root=true
type ReferenceSchema struct {
	metav1.TypeMeta `json:",inline"`

	// References is a list of reference paths.
	References []ReferencePath `json:"references,omitempty"`
}

// ReferencePath is a path to a reference in the XRD schema.
type ReferencePath struct {
	// JSONPath is the path to the reference in the XRD schema. It is a JSONPath
	// expression following RFC 9535 (https://tools.ietf.org/html/rfc6901) with
	// only name, non-negative index, and wildcard selector segments.
	// Examples: .spec.foo, .spec.foo[2], .spec.foo[*], .spec.foo[*].bar.
	// +required
	// +kubebuilder:validation:MinLength=1
	JSONPath string `json:"jsonPath"`

	// Kinds is a list of kinds that the reference can point to. If empty, the
	// reference can point to any kind.
	Kinds []ReferencableKind `json:"kinds,omitempty"`
}

// ReferencableKind is a referencable kind.
type ReferencableKind struct {
	// APIVersion is the API version of the kind.
	// +required
	// +kubebuilder:validation:MinLength=1
	APIVersion string `json:"apiVersion"`
	// Kind is the kind of the reference.
	// +required
	// +kubebuilder:validation:MinLength=1
	Kind string `json:"kind"`
}
