/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// VersionedPropertyDescriptor struct for VersionedPropertyDescriptor
type VersionedPropertyDescriptor struct {
	// The name of the property
	Name string `json:"name,omitempty"`
	// The display name of the property
	DisplayName string `json:"displayName,omitempty"`
	// Whether or not the property provides the identifier of a Controller Service
	IdentifiesControllerService bool `json:"identifiesControllerService,omitempty"`
	// Whether or not the property is considered sensitive
	Sensitive bool `json:"sensitive,omitempty"`
}
