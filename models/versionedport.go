/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// VersionedPort struct for VersionedPort
type VersionedPort struct {
	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`
	// The component's name
	Name string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments string   `json:"comments,omitempty"`
	Position Position `json:"position,omitempty"`
	// The type of port.
	Type string `json:"type,omitempty"`
	// The number of tasks that should be concurrently scheduled for the port.
	ConcurrentlySchedulableTaskCount int32 `json:"concurrentlySchedulableTaskCount,omitempty"`
	// The scheduled state of the component
	ScheduledState string `json:"scheduledState,omitempty"`
	// Whether or not this port allows remote access for site-to-site
	AllowRemoteAccess bool   `json:"allowRemoteAccess,omitempty"`
	ComponentType     string `json:"componentType,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`
}
