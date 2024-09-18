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
	"reflect"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	spacesv1alpha1 "github.com/upbound/up-sdk-go/apis/spaces/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Provider",type="string",JSONPath=".spec.objectStorage.provider"
// +kubebuilder:printcolumn:name="Bucket",type="string",JSONPath=".spec.objectStorage.bucket"
// +kubebuilder:printcolumn:name="Auth",type="string",JSONPath=".spec.objectStorage.credentials.source"
// +kubebuilder:printcolumn:name="Secret",type=string,JSONPath=`.metadata.annotations.spacebackupconfig\.admin\.internal\.spaces\.upbound\.io/secret`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories=spaces

// SpaceBackupConfig defines the configuration to backup a Space.
type SpaceBackupConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SpaceBackupConfigSpec `json:"spec"`
}

// +kubebuilder:object:root=true

// SpaceBackupConfigList contains a list of SpaceBackupConfig.
type SpaceBackupConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpaceBackupConfig `json:"items"`
}

// A SpaceBackupConfigSpec represents the configuration to backup or restore
// a Space.
type SpaceBackupConfigSpec struct {
	// ObjectStorage specifies the object storage configuration for the given provider.
	// +kubebuilder:validation:Required
	ObjectStorage SpaceBackupObjectStorage `json:"objectStorage"`
}

// SpaceBackupObjectStorage specifies the object storage configuration for the given provider.
type SpaceBackupObjectStorage struct {
	spacesv1alpha1.BackupObjectStorage `json:",inline"`

	// Credentials specifies the credentials to access the object storage.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="self.source != 'Secret' || (has(self.secretRef) && has(self.secretRef.name) && has(self.secretRef.__namespace__))",message="secretRef.name and namespace must be set when source is Secret"
	// +kubebuilder:validation:XValidation:rule="self.source != 'Environment' || (has(self.env) && has(self.env.name))",message="env.name must be set when source is Environment"
	Credentials SpaceBackupCredentials `json:"credentials"`
}

// SpaceBackupCredentials specifies the credentials to access the object storage.
type SpaceBackupCredentials struct {
	// Source of the credentials.
	// Source "Secret" requires "get" permissions on the referenced Secret.
	// +kubebuilder:validation:Enum=Secret;InjectedIdentity;Environment
	Source xpv1.CredentialsSource `json:"source"`

	// CommonCredentialSelectors provides common selectors for extracting
	// credentials.
	xpv1.CommonCredentialSelectors `json:",inline"`
}

var (
	// SpaceBackupConfigKind is the kind of a SpaceBackupConfig.
	SpaceBackupConfigKind = reflect.TypeOf(SpaceBackupConfig{}).Name()
)

func init() {
	SchemeBuilder.Register(&SpaceBackupConfig{}, &SpaceBackupConfigList{})
}
