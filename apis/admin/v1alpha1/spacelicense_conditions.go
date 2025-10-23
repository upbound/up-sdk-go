/*
Copyright 2025 The Upbound Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

// Condition types for SpaceLicense resource.
const (
	// TypeLicenseValid indicates whether the license is valid.
	TypeLicenseValid xpv1.ConditionType = "LicenseValid"

	// TypeUsageCompliant indicates whether the current usage is within the licensed capacity.
	TypeUsageCompliant xpv1.ConditionType = "UsageCompliant"

	// TypeMeasurementSucceeded indicates whether fetching measurements succeeded.
	TypeMeasurementSucceeded xpv1.ConditionType = "MeasurementSucceeded"
)

// Reasons for LicenseValid condition.
const (
	ReasonSignatureVerified     xpv1.ConditionReason = "SignatureVerified"
	ReasonFailedToGetLicenseKey xpv1.ConditionReason = "FailedToGetLicenseKey"
	ReasonLicenseExpiredInGrace xpv1.ConditionReason = "LicenseExpiredInGrace"
	ReasonLicenseExpiredFinal   xpv1.ConditionReason = "LicenseExpiredFinal"
	ReasonLicenseInvalid        xpv1.ConditionReason = "LicenseInvalid"
	ReasonCommunityEdition      xpv1.ConditionReason = "CommunityEdition"
	ReasonMeasurementActive     xpv1.ConditionReason = "MeasurementActive"
	ReasonMeasurementFailed     xpv1.ConditionReason = "MeasurementFailed"
)

// Reasons for UsageCompliant condition.
const (
	ReasonWithinCapacity  xpv1.ConditionReason = "WithinCapacity"
	ReasonExceedsCapacity xpv1.ConditionReason = "ExceedsCapacity"
)

// SpaceLicenseValid indicates that the license signature has been successfully verified.
func SpaceLicenseValid() xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeLicenseValid,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSignatureVerified,
		Message:            "The license signature has been successfully verified.",
	}
}

// SpaceLicenseCommunityEdition indicates that the license is active for community edition.
func SpaceLicenseCommunityEdition() xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeLicenseValid,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonCommunityEdition,
		Message:            "Community edition license is active.",
	}
}

// SpaceLicenseKeyGetFailed indicates that getting the license key failed.
// Use WithMessage() to provide specific details about what went wrong.
func SpaceLicenseKeyGetFailed() xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeLicenseValid,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonFailedToGetLicenseKey,
	}
}

// SpaceLicenseExpiredInGrace indicates that the license has expired but is still in grace period.
func SpaceLicenseExpiredInGrace(gracePeriodEnd time.Time) xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeLicenseValid,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonLicenseExpiredInGrace,
		Message:            "License has expired but is still functional until grace period ends on " + gracePeriodEnd.Format("2006-01-02") + ". Please renew urgently.",
	}
}

// SpaceLicenseExpiredFinal indicates that the license has expired and grace period has ended.
func SpaceLicenseExpiredFinal() xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeLicenseValid,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonLicenseExpiredFinal,
		Message:            "License has expired and grace period has ended. License is no longer functional.",
	}
}

// SpaceLicenseInvalid indicates that the license validation failed.
func SpaceLicenseInvalid(err error) xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeLicenseValid,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonLicenseInvalid,
		Message:            "License validation failed: " + err.Error(),
	}
}

// SpaceLicenseUsageCompliant indicates that the current usage is within the licensed capacity.
func SpaceLicenseUsageCompliant() xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeUsageCompliant,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonWithinCapacity,
		Message:            "Current usage is within the licensed capacity.",
	}
}

// SpaceLicenseUsageExceedsCapacity indicates that the current usage exceeds the licensed capacity.
func SpaceLicenseUsageExceedsCapacity(message string) xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeUsageCompliant,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonExceedsCapacity,
		Message:            message,
	}
}

// SpaceLicenseMeasurementSucceeded indicates that usage measurements are being collected successfully.
func SpaceLicenseMeasurementSucceeded() xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeMeasurementSucceeded,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonMeasurementActive,
		Message:            "Usage measurements are being collected successfully.",
	}
}

// SpaceLicenseMeasurementFailed indicates that fetching or processing measurements failed.
func SpaceLicenseMeasurementFailed(message string) xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeMeasurementSucceeded,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonMeasurementFailed,
		Message:            message,
	}
}
