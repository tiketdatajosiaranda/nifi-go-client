/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// FlowFileSummaryDto struct for FlowFileSummaryDto
type FlowFileSummaryDto struct {
	// The URI that can be used to access this FlowFile.
	Uri string `json:"uri,omitempty"`
	// The FlowFile UUID.
	Uuid string `json:"uuid,omitempty"`
	// The FlowFile filename.
	Filename string `json:"filename,omitempty"`
	// The FlowFile's position in the queue.
	Position int32 `json:"position,omitempty"`
	// The FlowFile file size.
	Size int64 `json:"size,omitempty"`
	// How long this FlowFile has been enqueued.
	QueuedDuration int64 `json:"queuedDuration,omitempty"`
	// Duration since the FlowFile's greatest ancestor entered the flow.
	LineageDuration int64 `json:"lineageDuration,omitempty"`
	// How long in milliseconds until the FlowFile penalty expires.
	PenaltyExpiresIn int64 `json:"penaltyExpiresIn,omitempty"`
	// The id of the node where this FlowFile resides.
	ClusterNodeId string `json:"clusterNodeId,omitempty"`
	// The label for the node where this FlowFile resides.
	ClusterNodeAddress string `json:"clusterNodeAddress,omitempty"`
	// If the FlowFile is penalized.
	Penalized bool `json:"penalized,omitempty"`
}
