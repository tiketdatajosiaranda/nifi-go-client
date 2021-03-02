/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ListingRequestDto struct for ListingRequestDto
type ListingRequestDto struct {
	// The id for this listing request.
	Id string `json:"id,omitempty"`
	// The URI for future requests to this listing request.
	Uri string `json:"uri,omitempty"`
	// The timestamp when the query was submitted.
	SubmissionTime string `json:"submissionTime,omitempty"`
	// The last time this listing request was updated.
	LastUpdated string `json:"lastUpdated,omitempty"`
	// The current percent complete.
	PercentCompleted int32 `json:"percentCompleted,omitempty"`
	// Whether the query has finished.
	Finished bool `json:"finished,omitempty"`
	// The reason, if any, that this listing request failed.
	FailureReason string `json:"failureReason,omitempty"`
	// The maximum number of FlowFileSummary objects to return
	MaxResults int32 `json:"maxResults,omitempty"`
	// The current state of the listing request.
	State     string       `json:"state,omitempty"`
	QueueSize QueueSizeDto `json:"queueSize,omitempty"`
	// The FlowFile summaries. The summaries will be populated once the request has completed.
	FlowFileSummaries []FlowFileSummaryDto `json:"flowFileSummaries,omitempty"`
	// Whether the source of the connection is running
	SourceRunning bool `json:"sourceRunning,omitempty"`
	// Whether the destination of the connection is running
	DestinationRunning bool `json:"destinationRunning,omitempty"`
}
