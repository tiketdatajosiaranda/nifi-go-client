/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// RemoteProcessGroupStatusSnapshotDto struct for RemoteProcessGroupStatusSnapshotDto
type RemoteProcessGroupStatusSnapshotDto struct {
	// The id of the remote process group.
	Id string `json:"id,omitempty"`
	// The id of the parent process group the remote process group resides in.
	GroupId string `json:"groupId,omitempty"`
	// The name of the remote process group.
	Name string `json:"name,omitempty"`
	// The URI of the target system.
	TargetUri string `json:"targetUri,omitempty"`
	// The transmission status of the remote process group.
	TransmissionStatus string `json:"transmissionStatus,omitempty"`
	// The number of active threads for the remote process group.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
	// The number of FlowFiles sent to the remote process group in the last 5 minutes.
	FlowFilesSent int32 `json:"flowFilesSent,omitempty"`
	// The size of the FlowFiles sent to the remote process group in the last 5 minutes.
	BytesSent int64 `json:"bytesSent,omitempty"`
	// The count/size of the flowfiles sent to the remote process group in the last 5 minutes.
	Sent string `json:"sent,omitempty"`
	// The number of FlowFiles received from the remote process group in the last 5 minutes.
	FlowFilesReceived int32 `json:"flowFilesReceived,omitempty"`
	// The size of the FlowFiles received from the remote process group in the last 5 minutes.
	BytesReceived int64 `json:"bytesReceived,omitempty"`
	// The count/size of the flowfiles received from the remote process group in the last 5 minutes.
	Received string `json:"received,omitempty"`
}
