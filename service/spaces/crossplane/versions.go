// Copyright 2024 Upbound Inc
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

package crossplane

// SupportedVersionsFile represents the structure of supported_versions.yaml file
// which contains the supported versions of Crossplane for a given Spaces version.
type SupportedVersionsFile struct {
	// APIVersion is the version of the file format.
	APIVersion string `json:"apiVersion" yaml:"apiVersion"`
	// SupportedVersions is a list of supported versions of Crossplane for a given Spaces version.
	SupportVersions []VersionMatrix `json:"supportedVersions" yaml:"supportedVersions"`
}

// VersionMatrix represents the supported versions of Crossplane for a given
// Spaces version.
type VersionMatrix struct {
	// SpacesVersion is the version of Spaces.
	SpacesVersion string `json:"spacesVersion" yaml:"spacesVersion"`
	// CrossplaneVersions is a list of supported versions of Crossplane for a given Spaces version.
	CrossplaneVersions []CrossplaneVersion `json:"crossplaneVersions" yaml:"crossplaneVersions"`
}

// CrossplaneVersion represents the supported version of Crossplane.
type CrossplaneVersion struct { // nolint: golint // Dropping Crossplane prefix and calling it Version feels odd, version everywhere.
	Version string `json:"version" yaml:"version"`
}
