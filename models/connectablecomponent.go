/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ConnectableComponent struct for ConnectableComponent
type ConnectableComponent struct {
	// The id of the connectable component.
	Id string `json:"id"`
	// The type of component the connectable is.
	Type string `json:"type"`
	// The id of the group that the connectable component resides in
	GroupId string `json:"groupId"`
	// The name of the connectable component
	Name string `json:"name,omitempty"`
	// The comments for the connectable component.
	Comments string `json:"comments,omitempty"`
}
