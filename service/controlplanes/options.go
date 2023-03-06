// Copyright 2023 Upbound Inc
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

package controlplanes

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/upbound/up-sdk-go/service/common"
)

const (
	// ConfigurationIDParam is the name of the Configuration UUID query parameter.
	ConfigurationIDParam = "configurationId"
)

// ControlPlaneListOption modifies a control plane list request, specifically.
type ControlPlaneListOption common.ListOption

// WithConfiguration sets the configurationId to filter the list response.
func WithConfiguration(id uuid.UUID) ControlPlaneListOption { // nolint:interfacer
	return func(r *http.Request) {
		q := r.URL.Query()
		q.Add(ConfigurationIDParam, id.String())
		r.URL.RawQuery = q.Encode()
	}
}
