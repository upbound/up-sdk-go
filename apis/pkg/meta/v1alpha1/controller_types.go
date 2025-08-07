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

// Package v1alpha1 contains API Schema definitions for the meta.pkg.upbound.io v1alpha1 API group.
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	v1 "github.com/crossplane/crossplane/v2/apis/pkg/meta/v1"
)

// ControllerPackagingType defines an enum for the controller package type.
type ControllerPackagingType string

const (
	// ControllerPackagingTypeHelm represents a controller that is packaged as a
	// Helm chart.
	ControllerPackagingTypeHelm ControllerPackagingType = "Helm"
)

// A Controller is the description of an Upbound Controller package.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Controller struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ControllerSpec `json:"spec"`
}

// ControllerSpec specifies the configuration of a Controller.
//
// +k8s:deepcopy-gen=true
type ControllerSpec struct {
	// PackagingType is the type of the Controller.
	// +kube:validation:Enum=Helm
	// +kubebuilder:default=Helm
	PackagingType ControllerPackagingType `json:"packagingType"`
	// Helm specific configuration.
	Helm *HelmSpec `json:"helm,omitempty"`

	v1.MetaSpec `json:",inline"`
}

// HelmSpec specifies the configuration of a Controller that is packaged as a
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

// GetCrossplaneConstraints gets the Controller package's Crossplane version
// constraints.
func (c *Controller) GetCrossplaneConstraints() *v1.CrossplaneConstraints {
	return c.Spec.MetaSpec.Crossplane
}

// GetDependencies gets the Controller package's dependencies.
func (c *Controller) GetDependencies() []v1.Dependency {
	return c.Spec.MetaSpec.DependsOn
}

// GetCapabilities gets the Controller package's capabilities.
func (c *Controller) GetCapabilities() []string {
	return c.Spec.Capabilities
}
