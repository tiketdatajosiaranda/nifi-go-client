/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ControllerServiceApiDto struct for ControllerServiceApiDto
type ControllerServiceApiDto struct {
	// The fully qualified name of the service interface.
	Type   string    `json:"type,omitempty"`
	Bundle BundleDto `json:"bundle,omitempty"`
}
