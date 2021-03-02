/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ConnectionStatisticsDto struct for ConnectionStatisticsDto
type ConnectionStatisticsDto struct {
	// The ID of the connection
	Id string `json:"id,omitempty"`
	// The timestamp of when the stats were last refreshed
	StatsLastRefreshed string                          `json:"statsLastRefreshed,omitempty"`
	AggregateSnapshot  ConnectionStatisticsSnapshotDto `json:"aggregateSnapshot,omitempty"`
	// A list of status snapshots for each node
	NodeSnapshots []NodeConnectionStatisticsSnapshotDto `json:"nodeSnapshots,omitempty"`
}
