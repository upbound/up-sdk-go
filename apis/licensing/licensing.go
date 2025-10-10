/*
Copyright 2025 The Upbound Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package licensing contains Kubernetes API groups for UXP licensing functionality.
package licensing

import (
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/upbound/up-sdk-go/apis/licensing/v1alpha1"
)

// Package type metadata.
const (
	Group   = "licensing.upbound.io"
	Version = "v1alpha1"
)

// AddToScheme adds all resources to the provided scheme.
func AddToScheme(s *runtime.Scheme) error {
	// Add all versions to the scheme.
	return v1alpha1.AddToScheme(s)
}
