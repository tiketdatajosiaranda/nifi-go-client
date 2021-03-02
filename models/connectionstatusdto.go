/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ConnectionStatusDto struct for ConnectionStatusDto
type ConnectionStatusDto struct {
	// The ID of the connection
	Id string `json:"id,omitempty"`
	// The ID of the Process Group that the connection belongs to
	GroupId string `json:"groupId,omitempty"`
	// The name of the connection
	Name string `json:"name,omitempty"`
	// The timestamp of when the stats were last refreshed
	StatsLastRefreshed string `json:"statsLastRefreshed,omitempty"`
	// The ID of the source component
	SourceId string `json:"sourceId,omitempty"`
	// The name of the source component
	SourceName string `json:"sourceName,omitempty"`
	// The ID of the destination component
	DestinationId string `json:"destinationId,omitempty"`
	// The name of the destination component
	DestinationName   string                      `json:"destinationName,omitempty"`
	AggregateSnapshot ConnectionStatusSnapshotDto `json:"aggregateSnapshot,omitempty"`
	// A list of status snapshots for each node
	NodeSnapshots []NodeConnectionStatusSnapshotDto `json:"nodeSnapshots,omitempty"`
}
