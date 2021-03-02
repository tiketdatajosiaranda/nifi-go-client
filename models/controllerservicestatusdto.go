/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ControllerServiceStatusDto struct for ControllerServiceStatusDto
type ControllerServiceStatusDto struct {
	// The run status of this ControllerService
	RunStatus string `json:"runStatus,omitempty"`
	// Indicates whether the component is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the component is valid)
	ValidationStatus string `json:"validationStatus,omitempty"`
	// The number of active threads for the component.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
}
