/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// PortStatusSnapshotDto struct for PortStatusSnapshotDto
type PortStatusSnapshotDto struct {
	// The id of the port.
	Id string `json:"id,omitempty"`
	// The id of the parent process group of the port.
	GroupId string `json:"groupId,omitempty"`
	// The name of the port.
	Name string `json:"name,omitempty"`
	// The active thread count for the port.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
	// The number of FlowFiles that have been accepted in the last 5 minutes.
	FlowFilesIn int32 `json:"flowFilesIn,omitempty"`
	// The size of hte FlowFiles that have been accepted in the last 5 minutes.
	BytesIn int64 `json:"bytesIn,omitempty"`
	// The count/size of flowfiles that have been accepted in the last 5 minutes.
	Input string `json:"input,omitempty"`
	// The number of FlowFiles that have been processed in the last 5 minutes.
	FlowFilesOut int32 `json:"flowFilesOut,omitempty"`
	// The number of bytes that have been processed in the last 5 minutes.
	BytesOut int64 `json:"bytesOut,omitempty"`
	// The count/size of flowfiles that have been processed in the last 5 minutes.
	Output string `json:"output,omitempty"`
	// Whether the port has incoming or outgoing connections to a remote NiFi.
	Transmitting bool `json:"transmitting,omitempty"`
	// The run status of the port.
	RunStatus string `json:"runStatus,omitempty"`
}
