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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

const (
	// Group is the API Group for projects.
	Group = "meta.pkg.upbound.io"
	// Version is the API version for projects.
	Version = "v1beta1"
	// GroupVersion is the GroupVersion for projects.
	GroupVersion = Group + "/" + Version
	// AddOnKind is the kind of an AddOn.
	AddOnKind = "AddOn"
)

var (
	// SchemeGroupVersion is group version used to register these objects.
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme adds all registered types to scheme.
	AddToScheme = SchemeBuilder.AddToScheme

	// AddOnGroupVersionKind is the GroupVersionKind for the AddOn type.
	AddOnGroupVersionKind = SchemeGroupVersion.WithKind(AddOnKind)
)

func init() {
	SchemeBuilder.Register(&AddOn{})
}
