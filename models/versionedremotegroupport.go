/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// VersionedRemoteGroupPort struct for VersionedRemoteGroupPort
type VersionedRemoteGroupPort struct {
	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`
	// The component's name
	Name string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments string   `json:"comments,omitempty"`
	Position Position `json:"position,omitempty"`
	// The id of the remote process group that the port resides in.
	RemoteGroupId string `json:"remoteGroupId,omitempty"`
	// The number of task that may transmit flowfiles to the target port concurrently.
	ConcurrentlySchedulableTaskCount int32 `json:"concurrentlySchedulableTaskCount,omitempty"`
	// Whether the flowfiles are compressed when sent to the target port.
	UseCompression bool      `json:"useCompression,omitempty"`
	BatchSize      BatchSize `json:"batchSize,omitempty"`
	ComponentType  string    `json:"componentType,omitempty"`
	// The ID of the port on the target NiFi instance
	TargetId string `json:"targetId,omitempty"`
	// The scheduled state of the component
	ScheduledState string `json:"scheduledState,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`
}
