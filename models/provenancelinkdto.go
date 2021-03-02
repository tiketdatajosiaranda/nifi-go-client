/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProvenanceLinkDto struct for ProvenanceLinkDto
type ProvenanceLinkDto struct {
	// The source node id of the link.
	SourceId string `json:"sourceId,omitempty"`
	// The target node id of the link.
	TargetId string `json:"targetId,omitempty"`
	// The flowfile uuid that traversed the link.
	FlowFileUuid string `json:"flowFileUuid,omitempty"`
	// The timestamp of the link (based on the destination).
	Timestamp string `json:"timestamp,omitempty"`
	// The timestamp of this link in milliseconds.
	Millis int64 `json:"millis,omitempty"`
}
