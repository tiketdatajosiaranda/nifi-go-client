/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ClusterDto struct for ClusterDto
type ClusterDto struct {
	// The collection of nodes that are part of the cluster.
	Nodes []NodeDto `json:"nodes,omitempty"`
	// The timestamp the report was generated.
	Generated string `json:"generated,omitempty"`
}
