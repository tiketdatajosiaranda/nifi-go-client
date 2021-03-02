/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// DropRequestDto struct for DropRequestDto
type DropRequestDto struct {
	// The id for this drop request.
	Id string `json:"id,omitempty"`
	// The URI for future requests to this drop request.
	Uri string `json:"uri,omitempty"`
	// The timestamp when the query was submitted.
	SubmissionTime string `json:"submissionTime,omitempty"`
	// The last time this drop request was updated.
	LastUpdated string `json:"lastUpdated,omitempty"`
	// The current percent complete.
	PercentCompleted int32 `json:"percentCompleted,omitempty"`
	// Whether the query has finished.
	Finished bool `json:"finished,omitempty"`
	// The reason, if any, that this drop request failed.
	FailureReason string `json:"failureReason,omitempty"`
	// The number of flow files currently queued.
	CurrentCount int32 `json:"currentCount,omitempty"`
	// The size of flow files currently queued in bytes.
	CurrentSize int64 `json:"currentSize,omitempty"`
	// The count and size of flow files currently queued.
	Current string `json:"current,omitempty"`
	// The number of flow files to be dropped as a result of this request.
	OriginalCount int32 `json:"originalCount,omitempty"`
	// The size of flow files to be dropped as a result of this request in bytes.
	OriginalSize int64 `json:"originalSize,omitempty"`
	// The count and size of flow files to be dropped as a result of this request.
	Original string `json:"original,omitempty"`
	// The number of flow files that have been dropped thus far.
	DroppedCount int32 `json:"droppedCount,omitempty"`
	// The size of flow files that have been dropped thus far in bytes.
	DroppedSize int64 `json:"droppedSize,omitempty"`
	// The count and size of flow files that have been dropped thus far.
	Dropped string `json:"dropped,omitempty"`
	// The current state of the drop request.
	State string `json:"state,omitempty"`
}
