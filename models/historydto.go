/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// HistoryDto struct for HistoryDto
type HistoryDto struct {
	// The number of number of actions that matched the search criteria..
	Total int32 `json:"total,omitempty"`
	// The timestamp when the report was generated.
	LastRefreshed string `json:"lastRefreshed,omitempty"`
	// The actions.
	Actions []ActionEntity `json:"actions,omitempty"`
}
