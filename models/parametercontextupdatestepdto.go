/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ParameterContextUpdateStepDto struct for ParameterContextUpdateStepDto
type ParameterContextUpdateStepDto struct {
	// Explanation of what happens in this step
	Description string `json:"description,omitempty"`
	// Whether or not this step has completed
	Complete bool `json:"complete,omitempty"`
	// An explanation of why this step failed, or null if this step did not fail
	FailureReason string `json:"failureReason,omitempty"`
}
