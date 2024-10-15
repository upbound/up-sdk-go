// Copyright 2023 Upbound Inc
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

// Package apis contains Kubernetes API for the MXE provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	adminv1alpha1 "github.com/upbound/up-sdk-go/apis/admin/v1alpha1"
	connectv1alpha1 "github.com/upbound/up-sdk-go/apis/connect/v1alpha1"
	policyv1alpha1 "github.com/upbound/up-sdk-go/apis/policy/v1alpha1"
	queryv1alpha1 "github.com/upbound/up-sdk-go/apis/query/v1alpha1"
	queryv1alpha2 "github.com/upbound/up-sdk-go/apis/query/v1alpha2"
	spacesv1alpha1 "github.com/upbound/up-sdk-go/apis/spaces/v1alpha1"
	spacesv1beta1 "github.com/upbound/up-sdk-go/apis/spaces/v1beta1"
	upboundv1alpha1 "github.com/upbound/up-sdk-go/apis/upbound/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		spacesv1beta1.SchemeBuilder.AddToScheme,
		spacesv1alpha1.SchemeBuilder.AddToScheme,
		queryv1alpha2.SchemeBuilder.AddToScheme,
		queryv1alpha1.SchemeBuilder.AddToScheme,
		upboundv1alpha1.SchemeBuilder.AddToScheme,
		policyv1alpha1.SchemeBuilder.AddToScheme,
		adminv1alpha1.SchemeBuilder.AddToScheme,
		connectv1alpha1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
