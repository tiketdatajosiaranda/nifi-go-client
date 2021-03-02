/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProvenanceNodeDto struct for ProvenanceNodeDto
type ProvenanceNodeDto struct {
	// The id of the node.
	Id string `json:"id,omitempty"`
	// The uuid of the flowfile associated with the provenance event.
	FlowFileUuid string `json:"flowFileUuid,omitempty"`
	// The uuid of the parent flowfiles of the provenance event.
	ParentUuids []string `json:"parentUuids,omitempty"`
	// The uuid of the childrent flowfiles of the provenance event.
	ChildUuids []string `json:"childUuids,omitempty"`
	// The identifier of the node that this event/flowfile originated from.
	ClusterNodeIdentifier string `json:"clusterNodeIdentifier,omitempty"`
	// The type of the node.
	Type string `json:"type,omitempty"`
	// If the type is EVENT, this is the type of event.
	EventType string `json:"eventType,omitempty"`
	// The timestamp of the node in milliseconds.
	Millis int64 `json:"millis,omitempty"`
	// The timestamp of the node formatted.
	Timestamp string `json:"timestamp,omitempty"`
}