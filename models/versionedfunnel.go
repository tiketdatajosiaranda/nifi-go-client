/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// VersionedFunnel struct for VersionedFunnel
type VersionedFunnel struct {
	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`
	// The component's name
	Name string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments      string   `json:"comments,omitempty"`
	Position      Position `json:"position,omitempty"`
	ComponentType string   `json:"componentType,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`
}
