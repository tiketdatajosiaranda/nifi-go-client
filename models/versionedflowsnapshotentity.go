/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// VersionedFlowSnapshotEntity struct for VersionedFlowSnapshotEntity
type VersionedFlowSnapshotEntity struct {
	VersionedFlowSnapshot VersionedFlowSnapshot `json:"versionedFlowSnapshot,omitempty"`
	ProcessGroupRevision  RevisionDto           `json:"processGroupRevision,omitempty"`
	// The ID of the Registry that this flow belongs to
	RegistryId string `json:"registryId,omitempty"`
	// If the Process Group to be updated has a child or descendant Process Group that is also under Version Control, this specifies whether or not the contents of that child/descendant Process Group should be updated.
	UpdateDescendantVersionedFlows bool `json:"updateDescendantVersionedFlows,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
}
