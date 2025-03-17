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

// Generated from spaces/v1beta1/sharedbackupconfig_types.go by ../hack/duplicate_api_type.sh. DO NOT EDIT.

package v1alpha1

import (
	"reflect"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/upbound/up-sdk-go/apis/common"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Provider",type="string",JSONPath=".spec.objectStorage.provider"
// +kubebuilder:printcolumn:name="Bucket",type="string",JSONPath=".spec.objectStorage.bucket"
// +kubebuilder:printcolumn:name="Auth",type="string",JSONPath=".spec.objectStorage.credentials.source"
// +kubebuilder:printcolumn:name="Secret",type="string",JSONPath=".spec.objectStorage.credentials.secretRef.name"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories=spaces

// SharedBackupConfig defines the configuration to backup and restore ControlPlanes.
type SharedBackupConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SharedBackupConfigSpec `json:"spec"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// SharedBackupConfigList contains a list of SharedBackupConfig.
type SharedBackupConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedBackupConfig `json:"items"`
}

// A SharedBackupConfigSpec represents the configuration to backup or restore
// ControlPlanes.
type SharedBackupConfigSpec struct {
	// ObjectStorage specifies the object storage configuration for the given provider.
	// +kubebuilder:validation:Required
	ObjectStorage BackupObjectStorage `json:"objectStorage"`
}

// BackupObjectStorageProvider define the name of an object storage provider.
type BackupObjectStorageProvider string

const (
	// BackupObjectStorageProviderAWS is the AWS object storage provider.
	BackupObjectStorageProviderAWS BackupObjectStorageProvider = "AWS"

	// BackupObjectStorageProviderAzure is the Azure object storage provider.
	BackupObjectStorageProviderAzure BackupObjectStorageProvider = "Azure"

	// BackupObjectStorageProviderGCP is the GCP object storage provider.
	BackupObjectStorageProviderGCP BackupObjectStorageProvider = "GCP"
)

// BackupObjectStorage specifies the object storage configuration for the given provider.
//
// +kubebuilder:validation:XValidation:rule="self.credentials.source != 'Secret' || (has(self.credentials.secretRef) && has(self.credentials.secretRef.name))",message="credentials.secretRef.name must be set when source is Secret"
type BackupObjectStorage struct {
	// Provider is the name of the object storage provider.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=AWS;Azure;GCP
	Provider BackupObjectStorageProvider `json:"provider"`

	// Bucket is the name of the bucket to store backups in.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Bucket string `json:"bucket"`

	// Prefix is the prefix to use for all backups using this
	// SharedBackupConfig, e.g. "prod/cluster1", resulting in backups for
	// controlplane "ctp1" in namespace "ns1" being stored in
	// "prod/cluster1/ns1/ctp1".
	Prefix string `json:"prefix,omitempty"`

	// Config is a free-form map of configuration options for the object storage provider.
	// See https://github.com/thanos-io/objstore?tab=readme-ov-file for more
	// information on the formats for each supported cloud provider. Bucket and
	// Provider will override the required values in the config.
	// +kubebuilder:pruning:PreserveUnknownFields
	Config common.JSONObject `json:"config,omitempty"`

	// Credentials specifies the credentials to access the object storage.
	// +kubebuilder:validation:Required
	Credentials BackupCredentials `json:"credentials"`
}

// BackupCredentials specifies the credentials to access the object storage.
type BackupCredentials struct {
	// Source of the credentials.
	// Source "Secret" requires "get" permissions on the referenced Secret.
	// +kubebuilder:validation:Enum=Secret;InjectedIdentity
	// +kubebuilder:validation:Required
	Source xpv1.CredentialsSource `json:"source"`

	// CommonCredentialSelectors provides common selectors for extracting
	// credentials.
	LocalCommonCredentialSelectors `json:",inline"`
}

// LocalCommonCredentialSelectors provides common selectors for extracting
// credentials.
type LocalCommonCredentialSelectors struct {
	// A SecretRef is a reference to a secret key that contains the credentials
	// that must be used to connect to the provider.
	// +optional
	SecretRef *LocalSecretKeySelector `json:"secretRef,omitempty"`
}

// A LocalSecretKeySelector is a reference to a secret key in a predefined
// namespace.
type LocalSecretKeySelector struct {
	xpv1.LocalSecretReference `json:",inline"`

	// The key to select.
	// +kubebuilder:default=credentials
	Key string `json:"key"`
}

var (
	// SharedBackupConfigKind is the kind of a SharedBackupConfig.
	SharedBackupConfigKind = reflect.TypeOf(SharedBackupConfig{}).Name()
)

func init() {
	SchemeBuilder.Register(&SharedBackupConfig{}, &SharedBackupConfigList{})
}
