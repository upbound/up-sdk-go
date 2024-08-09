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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// Condition types.
const (
	// TypeAcceptingChanges indicates the simulated control plane is accepting
	// changes. All changes detected while this condition is true will appear in
	// the change summary.
	TypeAcceptingChanges xpv1.ConditionType = "AcceptingChanges"
)

// Reasons a simulation is or isn't accepting changes.
const (
	ReasonSimulationStarting   xpv1.ConditionReason = "SimulationStarting"
	ReasonSimulationRunning    xpv1.ConditionReason = "SimulationRunning"
	ReasonSimulationComplete   xpv1.ConditionReason = "SimulationComplete"
	ReasonSimulationTerminated xpv1.ConditionReason = "SimulationTerminated"
)

// SimulationStarting indicates a Simulation is still starting. It isn't yet
// ready to accept changes.
func SimulationStarting() xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeAcceptingChanges,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSimulationStarting,
	}
}

// SimulationRunning indicates a Simulation is running, and ready to accept
// changes.
func SimulationRunning() xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeAcceptingChanges,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSimulationRunning,
	}
}

// SimulationComplete indicates a Simulation is complete, and no longer ready to
// accept changes. A SimulationComplete Simuluation's simulated control plane
// keeps running until the Simulation is terminated.
func SimulationComplete() xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeAcceptingChanges,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSimulationComplete,
	}
}

// SimulationTerminated indicates a Simulation is terminated. The simulated
// control plane is deleted when a simulation is terminated.
func SimulationTerminated() xpv1.Condition {
	return xpv1.Condition{
		Type:               TypeAcceptingChanges,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonSimulationTerminated,
	}
}
