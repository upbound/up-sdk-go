// Copyright 2024 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spaces

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

// Package type metadata.
const (
	// Group is the upbound group.
	Group = "upbound.io"
	// Version is the upbound api version.
	Version = "v1alpha1"
)

var (
	schemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}
	scheme             = runtime.NewScheme()
	codecs             = serializer.NewCodecFactory(scheme)
	parameterCodec     = runtime.NewParameterCodec(scheme)
	jsonSerializer     = json.NewSerializer(json.DefaultMetaFactory, scheme, scheme, false)
	codec              = codecs.CodecForVersions(jsonSerializer, jsonSerializer, schemeGroupVersion, schemeGroupVersion)
)

// AddToScheme adds the list of known types to scheme.
func AddToScheme(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(schemeGroupVersion,
		&Space{},
		&SpaceList{},
	)
	metav1.AddToGroupVersion(scheme, schemeGroupVersion)
	return nil
}

func init() {
	metav1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
