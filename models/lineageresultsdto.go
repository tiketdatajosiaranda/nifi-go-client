/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// LineageResultsDto struct for LineageResultsDto
type LineageResultsDto struct {
	// Any errors that occurred while generating the lineage.
	Errors []string `json:"errors,omitempty"`
	// The nodes in the lineage.
	Nodes []ProvenanceNodeDto `json:"nodes,omitempty"`
	// The links between the nodes in the lineage.
	Links []ProvenanceLinkDto `json:"links,omitempty"`
}
