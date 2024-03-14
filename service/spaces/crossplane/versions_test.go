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

import (
	_ "embed"
	"testing"

	"github.com/Masterminds/semver/v3"
	"gopkg.in/yaml.v3"
)

//go:embed supported_versions.yaml
var supportedVersionsFile string

func TestVersionFileValid(t *testing.T) {
	f := &SupportedVersionsFile{}
	if err := yaml.Unmarshal([]byte(supportedVersionsFile), f); err != nil {
		t.Fatalf("cannot parse supported_versions.yaml: %v", err)
	}
	if f.APIVersion != "v1" {
		t.Fatalf("expected APIVersion v1alpha1, got %s", f.APIVersion)
	}
	if len(f.SupportVersions) == 0 {
		t.Fatalf("expected non-empty SupportVersions")
	}
	for _, v := range f.SupportVersions {
		_, err := semver.NewVersion(v.SpacesVersion)
		if err != nil {
			t.Fatalf("expected valid semver version for SpacesVersion, got %s", v.SpacesVersion)
		}
		if len(v.CrossplaneVersions) == 0 {
			t.Fatalf("expected non-empty CrossplaneVersions")
		}
		for _, cv := range v.CrossplaneVersions {
			_, err = semver.NewVersion(cv)
			if err != nil {
				t.Fatalf("expected valid semver version for CrossplaneVersions, got %s", cv)
			}
		}
	}
}
