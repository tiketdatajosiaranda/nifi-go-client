/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProvenanceSearchableFieldDto struct for ProvenanceSearchableFieldDto
type ProvenanceSearchableFieldDto struct {
	// The id of the searchable field.
	Id string `json:"id,omitempty"`
	// The searchable field.
	Field string `json:"field,omitempty"`
	// The label for the searchable field.
	Label string `json:"label,omitempty"`
	// The type of the searchable field.
	Type string `json:"type,omitempty"`
}
