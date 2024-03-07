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
	ConditionTypeSourceSynced xpcommonv1.ConditionType   = "SourceSynced"
	ReasonSourceCompleted     xpcommonv1.ConditionReason = "Completed"
	ReasonSourceInProgress    xpcommonv1.ConditionReason = "InProgress"

	ConditionTypeScheduled xpcommonv1.ConditionType   = "Scheduled"
	ReasonScheduled        xpcommonv1.ConditionReason = "Scheduled"
	ReasonSchedulingError  xpcommonv1.ConditionReason = "SchedulingError"
	ReasonSchedulingFailed xpcommonv1.ConditionReason = "ScheduleFailed"
	ReasonDeploymentFailed xpcommonv1.ConditionReason = "DeploymentFailed"

	ConditionTypeSupported xpcommonv1.ConditionType   = "Supported"
	ReasonSupported        xpcommonv1.ConditionReason = "SupportedCrossplaneVersion"
	ReasonUnsupported      xpcommonv1.ConditionReason = "UnsupportedCrossplaneVersion"
)

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

// Scheduled returns a condition that indicates that scheduling of the
// control plane has succeeded.
func Scheduled() xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeScheduled,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonScheduled,
	}
}

// SchedulingError returns a condition that indicates that scheduling of the
// control plane had an error.
func SchedulingError(err error) xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeScheduled,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSchedulingError,
		Message:            err.Error(),
	}
}

// SchedulingFailed returns a condition that indicates that scheduling of the
// control plane did not succeed.
func SchedulingFailed(reason string) xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeScheduled,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSchedulingFailed,
		Message:            reason,
	}
}

// DeploymentFailed returns a condition that indicates that deployment of the
// control plane to a host cluster did not succeed.
func DeploymentFailed(err error) xpcommonv1.Condition {
	return xpcommonv1.Condition{
		Type:               ConditionTypeScheduled,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonDeploymentFailed,
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
