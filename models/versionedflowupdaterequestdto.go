/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// VersionedFlowUpdateRequestDto struct for VersionedFlowUpdateRequestDto
type VersionedFlowUpdateRequestDto struct {
	// The unique ID of this request.
	RequestId string `json:"requestId,omitempty"`
	// The unique ID of the Process Group being updated
	ProcessGroupId string `json:"processGroupId,omitempty"`
	// The URI for future requests to this drop request.
	Uri string `json:"uri,omitempty"`
	// The last time this request was updated.
	LastUpdated string `json:"lastUpdated,omitempty"`
	// Whether or not this request has completed
	Complete bool `json:"complete,omitempty"`
	// An explanation of why this request failed, or null if this request has not failed
	FailureReason string `json:"failureReason,omitempty"`
	// The percentage complete for the request, between 0 and 100
	PercentCompleted int32 `json:"percentCompleted,omitempty"`
	// The state of the request
	State                     string                       `json:"state,omitempty"`
	VersionControlInformation VersionControlInformationDto `json:"versionControlInformation,omitempty"`
}
