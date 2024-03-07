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
	// ConditionTypeSourceSynced indicates that the git source is in sync.
	ConditionTypeSourceSynced xpcommonv1.ConditionType = "SourceSynced"
	// ReasonSourceCompleted indicates that the git sync has been completed.
	ReasonSourceCompleted xpcommonv1.ConditionReason = "Completed"
	// ReasonSourceInProgress indicates that the git sync is still in progress.
	ReasonSourceInProgress xpcommonv1.ConditionReason = "InProgress"

	// ConditionTypeScheduled indicates that the control plane has been scheduled.
	ConditionTypeScheduled xpcommonv1.ConditionType = "Scheduled"
	// ReasonScheduled indicates that the control plane has been scheduled.
	ReasonScheduled xpcommonv1.ConditionReason = "Scheduled"
	// ReasonSchedulingError indicates that the control plane scheduling had an error.
	ReasonSchedulingError xpcommonv1.ConditionReason = "SchedulingError"
	// ReasonSchedulingFailed indicates that the control plane scheduling did not succeed
	// for non-error reasons, e.g. capacity.
	ReasonSchedulingFailed xpcommonv1.ConditionReason = "ScheduleFailed"
	// ReasonDeploymentFailed indicates that the control plane deployment did not succeed.
	ReasonDeploymentFailed xpcommonv1.ConditionReason = "DeploymentFailed"

	// ConditionTypeSupported indicates that the control plane is running a
	// supported version of Crossplane.
	ConditionTypeSupported xpcommonv1.ConditionType = "Supported"
	// ReasonSupported indicates that the control plane is running
	// a supported version of Crossplane.
	ReasonSupported xpcommonv1.ConditionReason = "SupportedCrossplaneVersion"
	// ReasonUnsupported indicates that the control plane is running a version
	// of Crossplane that is not supported.
	ReasonUnsupported xpcommonv1.ConditionReason = "UnsupportedCrossplaneVersion"
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
