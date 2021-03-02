/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// VariableDto struct for VariableDto
type VariableDto struct {
	// The name of the variable
	Name string `json:"name,omitempty"`
	// The value of the variable
	Value string `json:"value,omitempty"`
	// The ID of the Process Group where this Variable is defined
	ProcessGroupId string `json:"processGroupId,omitempty"`
	// A set of all components that will be affected if the value of this variable is changed
	AffectedComponents []AffectedComponentEntity `json:"affectedComponents,omitempty"`
}
