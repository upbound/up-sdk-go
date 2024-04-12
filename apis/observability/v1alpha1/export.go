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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/json"

	"github.com/upbound/up-sdk-go/apis/common"
)

// Export defines the telemetry exporter configuration.
// +kubebuilder:pruning:PreserveUnknownFields
type Export common.JSONObject

// OpenAPISchemaType is used by the kube-openapi generator when constructing
// the OpenAPI spec of this type.
//
// See: https://github.com/kubernetes/kube-openapi/tree/master/pkg/generators
func (j Export) OpenAPISchemaType() []string {
	// TODO: return actual types when anyOf is supported
	return nil
} // nolint:golint

// OpenAPISchemaFormat is used by the kube-openapi generator when constructing
// the OpenAPI spec of this type.
func (j Export) OpenAPISchemaFormat() string { return "" } // nolint:golint

// DeepCopy returns a deep copy of the Export.
func (j *Export) DeepCopy() *Export {
	if j == nil {
		return nil
	}
	return &Export{Object: runtime.DeepCopyJSONValue(j.Object).(map[string]interface{})}
}

// DeepCopyInto copies the receiver, writing into out.
func (j *Export) DeepCopyInto(target *Export) {
	if target == nil {
		return
	}
	if j == nil {
		target.Object = nil // shouldn't happen
		return
	}
	target.Object = runtime.DeepCopyJSONValue(j.Object).(map[string]interface{})
}

// MarshalJSON implements json.Marshaler.
func (j Export) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Object)
}

// UnmarshalJSON implements json.Marshaler.
func (j *Export) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &j.Object)
}

// String returns the JSON representation of the export.
func (j *Export) String() string {
	bs, _ := json.Marshal(j) // nolint:errcheck // no way to handle error here
	return string(bs)
}
