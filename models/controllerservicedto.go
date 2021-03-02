/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ControllerServiceDto struct for ControllerServiceDto
type ControllerServiceDto struct {
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string      `json:"parentGroupId,omitempty"`
	Position      PositionDto `json:"position,omitempty"`
	// The name of the controller service.
	Name string `json:"name,omitempty"`
	// The type of the controller service.
	Type   string    `json:"type,omitempty"`
	Bundle BundleDto `json:"bundle,omitempty"`
	// Lists the APIs this Controller Service implements.
	ControllerServiceApis []ControllerServiceApiDto `json:"controllerServiceApis,omitempty"`
	// The comments for the controller service.
	Comments string `json:"comments,omitempty"`
	// The state of the controller service.
	State string `json:"state,omitempty"`
	// Whether the controller service persists state.
	PersistsState bool `json:"persistsState,omitempty"`
	// Whether the controller service requires elevated privileges.
	Restricted bool `json:"restricted,omitempty"`
	// Whether the ontroller service has been deprecated.
	Deprecated bool `json:"deprecated,omitempty"`
	// Whether the controller service has multiple versions available.
	MultipleVersionsAvailable bool `json:"multipleVersionsAvailable,omitempty"`
	// The properties of the controller service.
	Properties map[string]string `json:"properties,omitempty"`
	// The descriptors for the controller service properties.
	Descriptors map[string]PropertyDescriptorDto `json:"descriptors,omitempty"`
	// The URL for the controller services custom configuration UI if applicable.
	CustomUiUrl string `json:"customUiUrl,omitempty"`
	// The annotation for the controller service. This is how the custom UI relays configuration to the controller service.
	AnnotationData string `json:"annotationData,omitempty"`
	// All components referencing this controller service.
	ReferencingComponents []ControllerServiceReferencingComponentEntity `json:"referencingComponents,omitempty"`
	// The validation errors from the controller service. These validation errors represent the problems with the controller service that must be resolved before it can be enabled.
	ValidationErrors []string `json:"validationErrors,omitempty"`
	// Indicates whether the ControllerService is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the ControllerService is valid)
	ValidationStatus string `json:"validationStatus,omitempty"`
	// Whether the underlying extension is missing.
	ExtensionMissing bool `json:"extensionMissing,omitempty"`
}