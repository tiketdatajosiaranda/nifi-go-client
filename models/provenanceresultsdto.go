/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProvenanceResultsDto struct for ProvenanceResultsDto
type ProvenanceResultsDto struct {
	// The provenance events that matched the search criteria.
	ProvenanceEvents []ProvenanceEventDto `json:"provenanceEvents,omitempty"`
	// The total number of results formatted.
	Total string `json:"total,omitempty"`
	// The total number of results.
	TotalCount int64 `json:"totalCount,omitempty"`
	// Then the search was performed.
	Generated string `json:"generated,omitempty"`
	// The oldest event available in the provenance repository.
	OldestEvent string `json:"oldestEvent,omitempty"`
	// The time offset of the server that's used for event time.
	TimeOffset int32 `json:"timeOffset,omitempty"`
	// Any errors that occurred while performing the provenance request.
	Errors []string `json:"errors,omitempty"`
}
