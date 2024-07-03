// Copyright 2023 Upbound Inc.
// All rights reserved

package v1beta1

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpcommonv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

const (
	// ConditionTypeHealthy indicates that the control plane is healthy.
	ConditionTypeHealthy xpcommonv1.ConditionType = "Healthy"
	// ReasonHealthy indicates that the control plane is healthy.
	ReasonHealthy xpcommonv1.ConditionReason = "HealthyControlPlane"
	// ReasonUnhealthy indicates that the control plane is unhealthy.
	ReasonUnhealthy xpcommonv1.ConditionReason = "UnhealthyControlPlane"

	// ConditionTypeSourceSynced indicates that the git source is in sync.
	ConditionTypeSourceSynced xpcommonv1.ConditionType = "SourceSynced"
	// ReasonSourceCompleted indicates that the git sync has been completed.
	ReasonSourceCompleted xpcommonv1.ConditionReason = "Completed"
	// ReasonSourceInProgress indicates that the git sync is still in progress.
	ReasonSourceInProgress xpcommonv1.ConditionReason = "InProgress"

	// ConditionTypeSupported indicates that the control plane is running a
	// supported version of Crossplane.
	ConditionTypeSupported xpcommonv1.ConditionType = "Supported"
	// ReasonSupported indicates that the control plane is running
	// a supported version of Crossplane.
	ReasonSupported xpcommonv1.ConditionReason = "SupportedCrossplaneVersion"
	// ReasonUnsupported indicates that the control plane is running a version
	// of Crossplane that is not supported.
	ReasonUnsupported xpcommonv1.ConditionReason = "UnsupportedCrossplaneVersion"

	// ConditionTypeRestored indicates that the control plane has been restored from backup.
	ConditionTypeRestored xpcommonv1.ConditionType = "Restored"
	// ReasonRestoreCompleted indicates that the control plane has been successfully restored from backup.
	ReasonRestoreCompleted xpcommonv1.ConditionReason = "Completed"
	// ReasonRestoreFailed indicates that the control plane failed to restore from backup.
	ReasonRestoreFailed xpcommonv1.ConditionReason = "Failed"

	// ReasonRestorePending indicates that the control plane restore is pending.
	ReasonRestorePending xpcommonv1.ConditionReason = "RestorePending"
)

func Healthy() xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeHealthy,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonHealthy,
	}
}

func Unhealthy(err error) xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeHealthy,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonUnhealthy,
		Message:            err.Error(),
	}
}

// SourceSynced returns a condition that indicates the control plane is in sync
// with the source.
func SourceSynced(revision string) xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeSourceSynced,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSourceCompleted,
		Message:            fmt.Sprintf("In sync with the revision %s", revision),
	}
}

// SourceInProgress returns a condition that indicates the control plane is still
// processing resources coming from the source.
func SourceInProgress(revision string) xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeSourceSynced,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSourceInProgress,
		Message:            fmt.Sprintf("Syncing revision %s", revision),
	}
}

// SourceError returns a condition that indicates the source operation of the
// control plane has failed.
func SourceError(err error) xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeSourceSynced,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSourceInProgress,
		Message:            err.Error(),
	}
}

// SupportedCrossplaneVersion returns a condition that indicates that the
// control plane is running a supported version of Crossplane.
func SupportedCrossplaneVersion() xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeSupported,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSupported,
	}
}

// UnsupportedCrossplaneVersion returns a condition that indicates that the
// control plane is running an unsupported version of Crossplane.
func UnsupportedCrossplaneVersion(msg string) xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeSupported,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonUnsupported,
		Message:            msg,
	}
}

// RestoreCompleted returns a condition that indicates that the control plane has been
// restored from backup.
func RestoreCompleted() xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeRestored,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonRestoreCompleted,
		Message:            "Control plane has been restored from specified backup",
	}
}

// RestoreFailed returns a condition that indicates that the control plane failed
// to restore from backup.
func RestoreFailed(err error) xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeRestored,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonRestoreFailed,
		Message:            err.Error(),
	}
}

// RestorePending returns a condition that indicates that the control plane restore
// is pending.
func RestorePending() xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               xpcommonv1.TypeReady,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonRestorePending,
		Message:            "Control plane restore is pending",
	}
}
