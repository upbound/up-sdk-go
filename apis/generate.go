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

//go:build generate
// +build generate

// NOTE: See the below link for details on what is happening here.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

// Generate deepcopy methodsets and CRD manifests
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=./...

// Generate crossplane-runtime methodsets (resource.Claim, etc)
//go:generate go run -tags generate github.com/crossplane/crossplane-tools/cmd/angryjet generate-methodsets --header-file=../hack/boilerplate.go.txt ./...

// Add license headers to all files.
//go:generate go run -tags generate github.com/google/addlicense -v -ignore **/kodata/*.yaml -ignore **/testdata/*.yaml -ignore **/vendor/** -f ../LICENSE -c "Upbound Inc" . ../cmd ../internal

package apis

import (
	_ "github.com/crossplane/crossplane-tools/cmd/angryjet" //nolint:typecheck
	_ "github.com/google/addlicense"                        //nolint:typecheck
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"     //nolint:typecheck
)
