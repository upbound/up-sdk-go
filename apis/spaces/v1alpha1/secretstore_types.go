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
	esv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SharedSecretStore represents a shared SecretStore projected as ClusterSecretStore
// into matching ControlPlanes in the same namespace. Once projected into a ControlPlane,
// it can be referenced from ExternalSecret instances, as part of `storeRef` fields.
// The secret store configuration including referenced credential are not leaked into the
// ControlPlanes and in that sense can be called secure as they are invisible to the
// ControlPlane workloads.
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Provisioned",type=string,JSONPath=`.metadata.annotations.sharedsecretstores\.internal\.spaces\.upbound\.io/provisioned-total`
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories={externalsecrets},shortName=sss
type SharedSecretStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SharedSecretStoreSpec   `json:"spec,omitempty"`
	Status SharedSecretStoreStatus `json:"status,omitempty"`
}

// SharedSecretStoreList contains a list of SharedSecretStore.
//
// +kubebuilder:object:root=true
type SharedSecretStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedSecretStore `json:"items"`
}

// Objects return the list of items.
func (s *SharedSecretStoreList) Objects() []client.Object {
	var objs = make([]client.Object, len(s.Items))
	for i := range s.Items {
		objs[i] = &s.Items[i]
	}
	return objs
}

// SharedSecretStoreSpec defines the desired state of SecretStore.
//
// +kubebuilder:validation:XValidation:rule="has(self.secretStoreName) == has(oldSelf.secretStoreName)",message="secretStoreName is immutable"
type SharedSecretStoreSpec struct {
	// SecretStoreName is the name to use when creating secret stores within a control plane.
	// optional, if not set, SharedSecretStore name will be used.
	// When set, it is immutable.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="value is immutable"
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:MinLength=1
	// +optional
	SecretStoreName string `json:"secretStoreName,omitempty"`

	// The metadata of the secret store to be created.
	// +optional
	SecretStoreMetadata *ResourceMetadata `json:"secretStoreMetadata,omitempty"`

	// The store is projected only to control planes
	// matching the provided selector. Either names or a labelSelector must be specified.
	// +kubebuilder:validation:XValidation:rule="(has(self.labelSelectors) || has(self.names)) && (size(self.labelSelectors) > 0 || size(self.names) > 0)",message="either names or a labelSelector must be specified"
	ControlPlaneSelector ResourceSelector `json:"controlPlaneSelector"`

	// The projected secret store can be consumed
	// only within namespaces matching the provided selector.
	// Either names or a labelSelector must be specified.
	// +kubebuilder:validation:XValidation:rule="(has(self.labelSelectors) || has(self.names)) && (size(self.labelSelectors) > 0 || size(self.names) > 0)",message="either names or a labelSelector must be specified"
	NamespaceSelector ResourceSelector `json:"namespaceSelector"`

	// Used to configure the provider. Only one provider may be set.
	Provider esv1beta1.SecretStoreProvider `json:"provider"`

	// Used to configure http retries if failed.
	// +optional
	RetrySettings *esv1beta1.SecretStoreRetrySettings `json:"retrySettings,omitempty"`

	// Used to configure store refresh interval in seconds.
	// +optional
	RefreshInterval int `json:"refreshInterval,omitempty"`
}

// SharedSecretStoreStatus defines the observed state of the SecretStore.
type SharedSecretStoreStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	// We needed to introduce a common field to workaround
	// https://github.com/kubernetes/kubernetes/issues/117447
	// otherwise the initial idea was that each controller
	// just updates/remove its item in the bellow lists.

	// observed resource generation.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// List of provisioning failures.
	// +optional
	// +listType=map
	// +listMapKey=controlPlane
	Failed []SecretStoreProvisioningFailure `json:"failed,omitempty"`

	// List of successfully provisioned targets.
	// +optional
	// +listType=map
	// +listMapKey=controlPlane
	Provisioned []SecretStoreProvisioningSuccess `json:"provisioned,omitempty"`
}

// SecretStoreProvisioningFailure defines secret store provisioning failure.
type SecretStoreProvisioningFailure struct {
	// ControlPlane name where the failure occurred.
	ControlPlane string `json:"controlPlane"`

	// List of occurred conditions.
	// +optional
	Conditions []esv1beta1.SecretStoreStatusCondition `json:"conditions,omitempty"`
}

// SecretStoreProvisioningSuccess defines secret store provision success.
type SecretStoreProvisioningSuccess struct {
	// ControlPlane name where the secret store got projected
	ControlPlane string `json:"controlPlane"`
}

// ResourceMetadata defines metadata fields for created resource.
type ResourceMetadata struct {

	// Annotations that are set on projected resource.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Labels that are set on projected resource.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

// ControlPlaneSelector returns a function that can be used for checking
// if a given object matches the selector.
func (c *SharedSecretStore) ControlPlaneSelector() func(obj client.Object) (bool, error) {
	return func(obj client.Object) (bool, error) {
		return c.Spec.ControlPlaneSelector.Matches(obj)
	}
}

var (
	// SharedSecretStoreKind is the kind of the SharedSecretStore.
	SharedSecretStoreKind = reflect.TypeOf(SharedSecretStore{}).Name()
)

func init() {
	SchemeBuilder.Register(&SharedSecretStore{}, &SharedSecretStoreList{})
}
