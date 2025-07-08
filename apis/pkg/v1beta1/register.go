// Copyright 2025 Upbound Inc
// All rights reserved

package v1beta1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "pkg.upbound.io"
	Version = "v1beta1"
)

var (
	// SchemeGroupVersion is group version used to register these objects.
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme adds the types in this package to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// AddOn type metadata.
var (
	AddOnKind                          = reflect.TypeOf(AddOn{}).Name()
	AddOnGroupVersionKind              = SchemeGroupVersion.WithKind(AddOnKind)
	AddOnRevisionKind                  = reflect.TypeOf(AddOnRevision{}).Name()
	AddOnRevisionGroupVersionKind      = SchemeGroupVersion.WithKind(AddOnRevisionKind)
	AddOnRuntimeConfigKind             = reflect.TypeOf(AddOnRuntimeConfig{}).Name()
	AddOnRuntimeConfigGroupVersionKind = SchemeGroupVersion.WithKind(AddOnRuntimeConfigKind)
)

func init() {
	SchemeBuilder.Register(&AddOn{}, &AddOnList{})
	SchemeBuilder.Register(&AddOnRevision{}, &AddOnRevisionList{})
	SchemeBuilder.Register(&AddOnRuntimeConfig{}, &AddOnRuntimeConfigList{})
}
