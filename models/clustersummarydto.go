/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ClusterSummaryDto struct for ClusterSummaryDto
type ClusterSummaryDto struct {
	// When clustered, reports the number of nodes connected vs the number of nodes in the cluster.
	ConnectedNodes string `json:"connectedNodes,omitempty"`
	// The number of nodes that are currently connected to the cluster
	ConnectedNodeCount int32 `json:"connectedNodeCount,omitempty"`
	// The number of nodes in the cluster, regardless of whether or not they are connected
	TotalNodeCount int32 `json:"totalNodeCount,omitempty"`
	// Whether this NiFi instance is clustered.
	Clustered bool `json:"clustered,omitempty"`
	// Whether this NiFi instance is connected to a cluster.
	ConnectedToCluster bool `json:"connectedToCluster,omitempty"`
}
