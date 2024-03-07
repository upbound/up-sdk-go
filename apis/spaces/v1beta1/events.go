// Copyright 2023 Upbound Inc.
// All rights reserved

package v1beta1

import "github.com/crossplane/crossplane-runtime/pkg/event"

const (
	EventReasonSourceControllerFailed event.Reason = "SourceControllerFailed"
	EventReasonSourcePullFailed       event.Reason = "SourcePullFailed"
	EventReasonSourceReadFailed       event.Reason = "SourceReadFailed"
	EventReasonSourceParseFailed      event.Reason = "SourceParseFailed"
	EventReasonSourceApplyFailed      event.Reason = "SourceApplyFailed"
	EventReasonSourceApplySucceeded   event.Reason = "SourceApplySucceeded"
	EventReasonSourceApplyInProgress  event.Reason = "SourceApplyInProgress"
)
