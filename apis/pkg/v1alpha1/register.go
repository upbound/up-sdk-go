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

// Package v1alpha1 contains pkg.upbound.io APIs in version v1alpha1.
//
// +kubebuilder:object:generate=true
// +groupName=pkg.upbound.io
// +versionName=v1alpha1
package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "pkg.upbound.io"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects.
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme adds the types in this package to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// RemoteConfiguration type metadata.
var (
	RemoteConfigurationKind                     = reflect.TypeOf(RemoteConfiguration{}).Name()
	RemoteConfigurationGroupVersionKind         = SchemeGroupVersion.WithKind(RemoteConfigurationKind)
	RemoteConfigurationRevisionKind             = reflect.TypeOf(RemoteConfigurationRevision{}).Name()
	RemoteConfigurationRevisionGroupVersionKind = SchemeGroupVersion.WithKind(RemoteConfigurationRevisionKind)
)

// Controller type metadata.
var (
	ControllerKind                          = reflect.TypeOf(Controller{}).Name()
	ControllerGroupVersionKind              = SchemeGroupVersion.WithKind(ControllerKind)
	ControllerRevisionKind                  = reflect.TypeOf(ControllerRevision{}).Name()
	ControllerRevisionGroupVersionKind      = SchemeGroupVersion.WithKind(ControllerRevisionKind)
	ControllerRuntimeConfigKind             = reflect.TypeOf(ControllerRuntimeConfig{}).Name()
	ControllerRuntimeConfigGroupVersionKind = SchemeGroupVersion.WithKind(ControllerRuntimeConfigKind)
)

func init() {
	SchemeBuilder.Register(&RemoteConfiguration{}, &RemoteConfigurationList{})
	SchemeBuilder.Register(&RemoteConfigurationRevision{}, &RemoteConfigurationRevisionList{})
	SchemeBuilder.Register(&Controller{}, &ControllerList{})
	SchemeBuilder.Register(&ControllerRevision{}, &ControllerRevisionList{})
	SchemeBuilder.Register(&ControllerRuntimeConfig{}, &ControllerRuntimeConfigList{})
}
