/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// VariableRegistryDto struct for VariableRegistryDto
type VariableRegistryDto struct {
	// The variables that are available in this Variable Registry
	Variables []VariableEntity `json:"variables,omitempty"`
	// The UUID of the Process Group that this Variable Registry belongs to
	ProcessGroupId string `json:"processGroupId,omitempty"`
}
