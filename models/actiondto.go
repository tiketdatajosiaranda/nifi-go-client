/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ActionDto struct for ActionDto
type ActionDto struct {
	// The action id.
	Id int32 `json:"id,omitempty"`
	// The identity of the user that performed the action.
	UserIdentity string `json:"userIdentity,omitempty"`
	// The timestamp of the action.
	Timestamp string `json:"timestamp,omitempty"`
	// The id of the source component.
	SourceId string `json:"sourceId,omitempty"`
	// The name of the source component.
	SourceName string `json:"sourceName,omitempty"`
	// The type of the source component.
	SourceType       string                 `json:"sourceType,omitempty"`
	ComponentDetails map[string]interface{} `json:"componentDetails,omitempty"`
	// The operation that was performed.
	Operation     string                 `json:"operation,omitempty"`
	ActionDetails map[string]interface{} `json:"actionDetails,omitempty"`
}
