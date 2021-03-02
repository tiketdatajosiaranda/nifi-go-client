/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProvenanceDto struct for ProvenanceDto
type ProvenanceDto struct {
	// The id of the provenance query.
	Id string `json:"id,omitempty"`
	// The URI for this query. Used for obtaining/deleting the request at a later time
	Uri string `json:"uri,omitempty"`
	// The timestamp when the query was submitted.
	SubmissionTime string `json:"submissionTime,omitempty"`
	// The timestamp when the query will expire.
	Expiration string `json:"expiration,omitempty"`
	// The current percent complete.
	PercentCompleted int32 `json:"percentCompleted,omitempty"`
	// Whether the query has finished.
	Finished bool                 `json:"finished,omitempty"`
	Request  ProvenanceRequestDto `json:"request,omitempty"`
	Results  ProvenanceResultsDto `json:"results,omitempty"`
}
