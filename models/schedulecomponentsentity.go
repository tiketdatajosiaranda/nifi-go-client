/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ScheduleComponentsEntity struct for ScheduleComponentsEntity
type ScheduleComponentsEntity struct {
	// The id of the ProcessGroup
	Id string `json:"id,omitempty"`
	// The desired state of the descendant components
	State string `json:"state,omitempty"`
	// Optional components to schedule. If not specified, all authorized descendant components will be used.
	Components map[string]RevisionDto `json:"components,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
}
