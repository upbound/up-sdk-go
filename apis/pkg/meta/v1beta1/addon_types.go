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

// Package v1beta1 contains API Schema definitions for the meta.pkg.upbound.io v1beta1 API group.
package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	v1 "github.com/crossplane/crossplane/v2/apis/pkg/meta/v1"
)

// AddOnPackagingType defines an enum for the addon package type.
type AddOnPackagingType string

const (
	// AddOnPackagingTypeHelm represents an addon that is packaged as a
	// Helm chart.
	AddOnPackagingTypeHelm AddOnPackagingType = "Helm"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// An AddOn is the description of an Upbound AddOn package.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AddOn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AddOnSpec `json:"spec"`
}

// AddOnSpec specifies the configuration of an AddOn.
//
// +k8s:deepcopy-gen=true
type AddOnSpec struct {
	// PackagingType is the type of the AddOn.
	// +kube:validation:Enum=Helm
	// +kubebuilder:default=Helm
	PackagingType AddOnPackagingType `json:"packagingType"`
	// Helm specific configuration.
	Helm *HelmSpec `json:"helm,omitempty"`

	v1.MetaSpec `json:",inline"`
}

// HelmSpec specifies the configuration of an AddOn that is packaged as a
// Helm chart.
//
// +k8s:deepcopy-gen=true
type HelmSpec struct {
	// ReleaseName is the release name to install the Helm chart with.
	ReleaseName string `json:"releaseName"`
	// ReleaseNamespace is the release namespace to install the Helm chart into.
	ReleaseNamespace string `json:"releaseNamespace"`
	// Values is the values to be passed to the Helm chart.
	// +kubebuilder:pruning:PreserveUnknownFields
	Values runtime.RawExtension `json:"values,omitempty"`
}

// GetCrossplaneConstraints gets the AddOn package's Crossplane version
// constraints.
func (a *AddOn) GetCrossplaneConstraints() *v1.CrossplaneConstraints {
	return a.Spec.MetaSpec.Crossplane
}

// GetDependencies gets the AddOn package's dependencies.
func (a *AddOn) GetDependencies() []v1.Dependency {
	return a.Spec.MetaSpec.DependsOn
}

// GetCapabilities gets the AddOn package's capabilities.
func (a *AddOn) GetCapabilities() []string {
	return a.Spec.Capabilities
}
