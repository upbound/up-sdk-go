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

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

const (
	// ConditionSecretRefReady indicates the status of any referenced secrets
	// for API objects that reference a secret.
	ConditionSecretRefReady = xpv1.ConditionType("SecretRefsReady")

	// ReasonSecretRefMissing is added to object if the referenced secret(s) do not exist
	ReasonSecretRefMissing = xpv1.ConditionReason("ReferencedSecretNotFound")
	// ReasonSecretRefMissingKey is added when the object if the secret(s) exists, but the key data is missing
	ReasonSecretRefMissingKey = xpv1.ConditionReason("ReferencedSecretDataNotFound")
	// ReasonSecretRefReady is added when the referenced secret(s) and data are found
	ReasonSecretRefReady = xpv1.ConditionReason("ReferencedSecretRefReady")
	// ReasonSecretRefNone is added  when object not reference a secret
	ReasonSecretRefNone = xpv1.ConditionReason("NoReferencedSecret")
)
