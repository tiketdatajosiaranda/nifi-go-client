/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ConnectableDto struct for ConnectableDto
type ConnectableDto struct {
	// The id of the connectable component.
	Id string `json:"id"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
	// The type of component the connectable is.
	Type string `json:"type"`
	// The id of the group that the connectable component resides in
	GroupId string `json:"groupId"`
	// The name of the connectable component
	Name string `json:"name,omitempty"`
	// Reflects the current state of the connectable component.
	Running bool `json:"running,omitempty"`
	// If the connectable component represents a remote port, indicates if the target is configured to transmit.
	Transmitting bool `json:"transmitting,omitempty"`
	// If the connectable component represents a remote port, indicates if the target exists.
	Exists bool `json:"exists,omitempty"`
	// The comments for the connectable component.
	Comments string `json:"comments,omitempty"`
}
