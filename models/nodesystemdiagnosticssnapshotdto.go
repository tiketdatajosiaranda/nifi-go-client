/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// NodeSystemDiagnosticsSnapshotDto struct for NodeSystemDiagnosticsSnapshotDto
type NodeSystemDiagnosticsSnapshotDto struct {
	// The unique ID that identifies the node
	NodeId string `json:"nodeId,omitempty"`
	// The API address of the node
	Address string `json:"address,omitempty"`
	// The API port used to communicate with the node
	ApiPort  int32                        `json:"apiPort,omitempty"`
	Snapshot SystemDiagnosticsSnapshotDto `json:"snapshot,omitempty"`
}
