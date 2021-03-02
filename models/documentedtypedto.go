/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// DocumentedTypeDto struct for DocumentedTypeDto
type DocumentedTypeDto struct {
	// The fully qualified name of the type.
	Type   string    `json:"type,omitempty"`
	Bundle BundleDto `json:"bundle,omitempty"`
	// If this type represents a ControllerService, this lists the APIs it implements.
	ControllerServiceApis []ControllerServiceApiDto `json:"controllerServiceApis,omitempty"`
	// The description of the type.
	Description string `json:"description,omitempty"`
	// Whether this type is restricted.
	Restricted bool `json:"restricted,omitempty"`
	// The optional description of why the usage of this component is restricted.
	UsageRestriction string `json:"usageRestriction,omitempty"`
	// An optional collection of explicit restrictions. If specified, these explicit restrictions will be enfored.
	ExplicitRestrictions []ExplicitRestrictionDto `json:"explicitRestrictions,omitempty"`
	// The description of why the usage of this component is restricted.
	DeprecationReason string `json:"deprecationReason,omitempty"`
	// The tags associated with this type.
	Tags []string `json:"tags,omitempty"`
}
