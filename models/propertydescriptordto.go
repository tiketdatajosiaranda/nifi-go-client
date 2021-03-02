/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// PropertyDescriptorDto struct for PropertyDescriptorDto
type PropertyDescriptorDto struct {
	// The name for the property.
	Name string `json:"name,omitempty"`
	// The human readable name for the property.
	DisplayName string `json:"displayName,omitempty"`
	// The description for the property. Used to relay additional details to a user or provide a mechanism of documenting intent.
	Description string `json:"description,omitempty"`
	// The default value for the property.
	DefaultValue string `json:"defaultValue,omitempty"`
	// Allowable values for the property. If empty then the allowed values are not constrained.
	AllowableValues []AllowableValueEntity `json:"allowableValues,omitempty"`
	// Whether the property is required.
	Required bool `json:"required,omitempty"`
	// Whether the property is sensitive and protected whenever stored or represented.
	Sensitive bool `json:"sensitive,omitempty"`
	// Whether the property is dynamic (user-defined).
	Dynamic bool `json:"dynamic,omitempty"`
	// Whether the property supports expression language.
	SupportsEl bool `json:"supportsEl,omitempty"`
	// Scope of the Expression Language evaluation for the property.
	ExpressionLanguageScope string `json:"expressionLanguageScope,omitempty"`
	// If the property identifies a controller service this returns the fully qualified type.
	IdentifiesControllerService       string    `json:"identifiesControllerService,omitempty"`
	IdentifiesControllerServiceBundle BundleDto `json:"identifiesControllerServiceBundle,omitempty"`
}
