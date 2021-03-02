/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// LineageDto struct for LineageDto
type LineageDto struct {
	// The id of this lineage query.
	Id string `json:"id,omitempty"`
	// The URI for this lineage query for later retrieval and deletion.
	Uri string `json:"uri,omitempty"`
	// When the lineage query was submitted.
	SubmissionTime string `json:"submissionTime,omitempty"`
	// When the lineage query will expire.
	Expiration string `json:"expiration,omitempty"`
	// The percent complete for the lineage query.
	PercentCompleted int32 `json:"percentCompleted,omitempty"`
	// Whether the lineage query has finished.
	Finished bool              `json:"finished,omitempty"`
	Request  LineageRequestDto `json:"request,omitempty"`
	Results  LineageResultsDto `json:"results,omitempty"`
}
