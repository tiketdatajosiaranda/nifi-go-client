/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// RemoteProcessGroupPortDto struct for RemoteProcessGroupPortDto
type RemoteProcessGroupPortDto struct {
	// The id of the port.
	Id string `json:"id,omitempty"`
	// The id of the target port.
	TargetId string `json:"targetId,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
	// The id of the remote process group that the port resides in.
	GroupId string `json:"groupId,omitempty"`
	// The name of the target port.
	Name string `json:"name,omitempty"`
	// The comments as configured on the target port.
	Comments string `json:"comments,omitempty"`
	// The number of task that may transmit flowfiles to the target port concurrently.
	ConcurrentlySchedulableTaskCount int32 `json:"concurrentlySchedulableTaskCount,omitempty"`
	// Whether the remote port is configured for transmission.
	Transmitting bool `json:"transmitting,omitempty"`
	// Whether the flowfiles are compressed when sent to the target port.
	UseCompression bool `json:"useCompression,omitempty"`
	// Whether the target port exists.
	Exists bool `json:"exists,omitempty"`
	// Whether the target port is running.
	TargetRunning bool `json:"targetRunning,omitempty"`
	// Whether the port has either an incoming or outgoing connection.
	Connected     bool             `json:"connected,omitempty"`
	BatchSettings BatchSettingsDto `json:"batchSettings,omitempty"`
}