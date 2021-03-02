/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ControllerStatusDto struct for ControllerStatusDto
type ControllerStatusDto struct {
	// The number of active threads in the NiFi.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
	// The number of terminated threads in the NiFi.
	TerminatedThreadCount int32 `json:"terminatedThreadCount,omitempty"`
	// The number of flowfiles queued in the NiFi.
	Queued string `json:"queued,omitempty"`
	// The number of FlowFiles queued across the entire flow
	FlowFilesQueued int32 `json:"flowFilesQueued,omitempty"`
	// The size of the FlowFiles queued across the entire flow
	BytesQueued int64 `json:"bytesQueued,omitempty"`
	// The number of running components in the NiFi.
	RunningCount int32 `json:"runningCount,omitempty"`
	// The number of stopped components in the NiFi.
	StoppedCount int32 `json:"stoppedCount,omitempty"`
	// The number of invalid components in the NiFi.
	InvalidCount int32 `json:"invalidCount,omitempty"`
	// The number of disabled components in the NiFi.
	DisabledCount int32 `json:"disabledCount,omitempty"`
	// The number of active remote ports in the NiFi.
	ActiveRemotePortCount int32 `json:"activeRemotePortCount,omitempty"`
	// The number of inactive remote ports in the NiFi.
	InactiveRemotePortCount int32 `json:"inactiveRemotePortCount,omitempty"`
	// The number of up to date versioned process groups in the NiFi.
	UpToDateCount int32 `json:"upToDateCount,omitempty"`
	// The number of locally modified versioned process groups in the NiFi.
	LocallyModifiedCount int32 `json:"locallyModifiedCount,omitempty"`
	// The number of stale versioned process groups in the NiFi.
	StaleCount int32 `json:"staleCount,omitempty"`
	// The number of locally modified and stale versioned process groups in the NiFi.
	LocallyModifiedAndStaleCount int32 `json:"locallyModifiedAndStaleCount,omitempty"`
	// The number of versioned process groups in the NiFi that are unable to sync to a registry.
	SyncFailureCount int32 `json:"syncFailureCount,omitempty"`
}
