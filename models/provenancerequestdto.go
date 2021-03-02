/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProvenanceRequestDto struct for ProvenanceRequestDto
type ProvenanceRequestDto struct {
	// The search terms used to perform the search.
	SearchTerms map[string]string `json:"searchTerms,omitempty"`
	// The id of the node in the cluster where this provenance originated.
	ClusterNodeId string `json:"clusterNodeId,omitempty"`
	// The earliest event time to include in the query.
	StartDate string `json:"startDate,omitempty"`
	// The latest event time to include in the query.
	EndDate string `json:"endDate,omitempty"`
	// The minimum file size to include in the query.
	MinimumFileSize string `json:"minimumFileSize,omitempty"`
	// The maximum file size to include in the query.
	MaximumFileSize string `json:"maximumFileSize,omitempty"`
	// The maximum number of results to include.
	MaxResults int32 `json:"maxResults,omitempty"`
	// Whether or not to summarize provenance events returned. This property is false by default.
	Summarize bool `json:"summarize,omitempty"`
	// Whether or not incremental results are returned. If false, provenance events are only returned once the query completes. This property is true by default.
	IncrementalResults bool `json:"incrementalResults,omitempty"`
}
