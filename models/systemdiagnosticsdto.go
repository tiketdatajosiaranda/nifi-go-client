/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// SystemDiagnosticsDto struct for SystemDiagnosticsDto
type SystemDiagnosticsDto struct {
	AggregateSnapshot SystemDiagnosticsSnapshotDto `json:"aggregateSnapshot,omitempty"`
	// A systems diagnostics snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
	NodeSnapshots []NodeSystemDiagnosticsSnapshotDto `json:"nodeSnapshots,omitempty"`
}
