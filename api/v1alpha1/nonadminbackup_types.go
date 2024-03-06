/*
Copyright 2024.

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
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NonAdminBackupPhase is a simple one high-level summary of the lifecycle of an NonAdminBackup.
// +kubebuilder:validation:Enum=New;BackingOff;Created
type NonAdminBackupPhase string

const (
	// NonAdminBackupPhaseNew - NonAdminBackup resource was accepted by the OpenShift cluster, but it has not yet been processed by the NonAdminController
	NonAdminBackupPhaseNew NonAdminBackupPhase = "New"
	// NonAdminBackupPhaseBackingOff - Velero Backup object was not created due to NonAdminBackup error (configuration or similar)
	NonAdminBackupPhaseBackingOff NonAdminBackupPhase = "BackingOff"
	// NonAdminBackupPhaseCreated - Velero Backup was created. The Phase will not have additional informations about the Backup.
	NonAdminBackupPhaseCreated NonAdminBackupPhase = "Created"
)

// NonAdminCondition are used for more detailed information supporing NonAdminBackupPhase state.
// +kubebuilder:validation:Enum=BackupAccepted;BackupQueued
type NonAdminCondition string

// Predefined conditions for NonAdminBackup.
// One NonAdminBackup object may have multiple conditions.
// It is more granular knowledge of the NonAdminBackup object and represents the
// array of the conditions through which the NonAdminBackup has or has not passed
const (
	NonAdminBackupConditionAccepted NonAdminCondition = "BackupAccepted"
	NonAdminBackupConditionQueued   NonAdminCondition = "BackupQueued"
)

// NonAdminBackupSpec defines the desired state of NonAdminBackup
type NonAdminBackupSpec struct {
	// BackupSpec defines the specification for a Velero backup.
	BackupSpec *velerov1api.BackupSpec `json:"backupSpec,omitempty"`

	// NonAdminBackup log level (use debug for the most logging, leave unset for default)
	// +optional
	// +kubebuilder:validation:Enum=trace;debug;info;warning;error;fatal;panic
	LogLevel string `json:"logLevel,omitempty"`
}

// NonAdminBackupStatus defines the observed state of NonAdminBackup
type NonAdminBackupStatus struct {
	// OadpVeleroBackup references the VeleroBackup object. Format: <name>.<namespace>
	// +optional
	OadpVeleroBackup string `json:"oadpVeleroBackup,omitempty"`

	// BackupStatus captures the current status of a Velero backup.
	// +optional
	BackupStatus *velerov1api.BackupStatus `json:"backupStatus,omitempty"`

	Phase      NonAdminBackupPhase `json:"phase,omitempty"`
	Conditions []metav1.Condition  `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// NonAdminBackup is the Schema for the nonadminbackups API
type NonAdminBackup struct {
	Spec   NonAdminBackupSpec   `json:"spec,omitempty"`
	Status NonAdminBackupStatus `json:"status,omitempty"`

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

// +kubebuilder:object:root=true

// NonAdminBackupList contains a list of NonAdminBackup
type NonAdminBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NonAdminBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NonAdminBackup{}, &NonAdminBackupList{})
}
