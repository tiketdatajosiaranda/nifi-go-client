/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProcessGroupNameDto struct for ProcessGroupNameDto
type ProcessGroupNameDto struct {
	// The ID of the Process Group
	Id string `json:"id,omitempty"`
	// The name of the Process Group, or the ID of the Process Group if the user does not have the READ policy for the Process Group
	Name string `json:"name,omitempty"`
}
