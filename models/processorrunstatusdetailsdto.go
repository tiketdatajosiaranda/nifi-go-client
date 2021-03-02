/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProcessorRunStatusDetailsDto struct for ProcessorRunStatusDetailsDto
type ProcessorRunStatusDetailsDto struct {
	// The ID of the processor
	Id string `json:"id,omitempty"`
	// The name of the processor
	Name string `json:"name,omitempty"`
	// The run status of the processor
	RunStatus string `json:"runStatus,omitempty"`
	// The processor's validation errors
	ValidationErrors []string `json:"validationErrors,omitempty"`
	// The current number of threads that the processor is currently using
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
}
