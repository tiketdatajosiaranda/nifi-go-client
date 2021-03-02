/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// RemoteProcessGroupStatusDto struct for RemoteProcessGroupStatusDto
type RemoteProcessGroupStatusDto struct {
	// The unique ID of the process group that the Processor belongs to
	GroupId string `json:"groupId,omitempty"`
	// The unique ID of the Processor
	Id string `json:"id,omitempty"`
	// The name of the remote process group.
	Name string `json:"name,omitempty"`
	// The URI of the target system.
	TargetUri string `json:"targetUri,omitempty"`
	// The transmission status of the remote process group.
	TransmissionStatus string `json:"transmissionStatus,omitempty"`
	// The time the status for the process group was last refreshed.
	StatsLastRefreshed string `json:"statsLastRefreshed,omitempty"`
	// Indicates whether the component is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the component is valid)
	ValidationStatus  string                              `json:"validationStatus,omitempty"`
	AggregateSnapshot RemoteProcessGroupStatusSnapshotDto `json:"aggregateSnapshot,omitempty"`
	// A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
	NodeSnapshots []NodeRemoteProcessGroupStatusSnapshotDto `json:"nodeSnapshots,omitempty"`
}
