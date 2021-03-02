/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// StateMapDto struct for StateMapDto
type StateMapDto struct {
	// The scope of this StateMap.
	Scope string `json:"scope,omitempty"`
	// The total number of state entries. When the state map is lengthy, only of portion of the entries are returned.
	TotalEntryCount int32 `json:"totalEntryCount,omitempty"`
	// The state.
	State []StateEntryDto `json:"state,omitempty"`
}
