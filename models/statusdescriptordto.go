/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// StatusDescriptorDto struct for StatusDescriptorDto
type StatusDescriptorDto struct {
	// The name of the status field.
	Field string `json:"field,omitempty"`
	// The label for the status field.
	Label string `json:"label,omitempty"`
	// The description of the status field.
	Description string `json:"description,omitempty"`
	// The formatter for the status descriptor.
	Formatter string `json:"formatter,omitempty"`
}
