/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ControllerServiceReferencingComponentDto struct for ControllerServiceReferencingComponentDto
type ControllerServiceReferencingComponentDto struct {
	// The group id for the component referencing a controller service. If this component is another controller service or a reporting task, this field is blank.
	GroupId string `json:"groupId,omitempty"`
	// The id of the component referencing a controller service.
	Id string `json:"id,omitempty"`
	// The name of the component referencing a controller service.
	Name string `json:"name,omitempty"`
	// The type of the component referencing a controller service in simple Java class name format without package name.
	Type string `json:"type,omitempty"`
	// The scheduled state of a processor or reporting task referencing a controller service. If this component is another controller service, this field represents the controller service state.
	State string `json:"state,omitempty"`
	// The properties for the component.
	Properties map[string]string `json:"properties,omitempty"`
	// The descriptors for the component properties.
	Descriptors map[string]PropertyDescriptorDto `json:"descriptors,omitempty"`
	// The validation errors for the component.
	ValidationErrors []string `json:"validationErrors,omitempty"`
	// The type of reference this is.
	ReferenceType string `json:"referenceType,omitempty"`
	// The number of active threads for the referencing component.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
	// If the referencing component represents a controller service, this indicates whether it has already been represented in this hierarchy.
	ReferenceCycle bool `json:"referenceCycle,omitempty"`
	// If the referencing component represents a controller service, these are the components that reference it.
	ReferencingComponents []ControllerServiceReferencingComponentEntity `json:"referencingComponents,omitempty"`
}
