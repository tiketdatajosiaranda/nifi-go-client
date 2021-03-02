/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// RelationshipDto struct for RelationshipDto
type RelationshipDto struct {
	// The relationship name.
	Name string `json:"name,omitempty"`
	// The relationship description.
	Description string `json:"description,omitempty"`
	// Whether or not flowfiles sent to this relationship should auto terminate.
	AutoTerminate bool `json:"autoTerminate,omitempty"`
}
