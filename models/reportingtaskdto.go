/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ReportingTaskDto struct for ReportingTaskDto
type ReportingTaskDto struct {
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string      `json:"parentGroupId,omitempty"`
	Position      PositionDto `json:"position,omitempty"`
	// The name of the reporting task.
	Name string `json:"name,omitempty"`
	// The fully qualified type of the reporting task.
	Type   string    `json:"type,omitempty"`
	Bundle BundleDto `json:"bundle,omitempty"`
	// The state of the reporting task.
	State string `json:"state,omitempty"`
	// The comments of the reporting task.
	Comments string `json:"comments,omitempty"`
	// Whether the reporting task persists state.
	PersistsState bool `json:"persistsState,omitempty"`
	// Whether the reporting task requires elevated privileges.
	Restricted bool `json:"restricted,omitempty"`
	// Whether the reporting task has been deprecated.
	Deprecated bool `json:"deprecated,omitempty"`
	// Whether the reporting task has multiple versions available.
	MultipleVersionsAvailable bool `json:"multipleVersionsAvailable,omitempty"`
	// The frequency with which to schedule the reporting task. The format of the value willd epend on the valud of the schedulingStrategy.
	SchedulingPeriod string `json:"schedulingPeriod,omitempty"`
	// The scheduling strategy that determines how the schedulingPeriod value should be interpreted.
	SchedulingStrategy string `json:"schedulingStrategy,omitempty"`
	// The default scheduling period for the different scheduling strategies.
	DefaultSchedulingPeriod map[string]string `json:"defaultSchedulingPeriod,omitempty"`
	// The properties of the reporting task.
	Properties map[string]string `json:"properties,omitempty"`
	// The descriptors for the reporting tasks properties.
	Descriptors map[string]PropertyDescriptorDto `json:"descriptors,omitempty"`
	// The URL for the custom configuration UI for the reporting task.
	CustomUiUrl string `json:"customUiUrl,omitempty"`
	// The annotation data for the repoting task. This is how the custom UI relays configuration to the reporting task.
	AnnotationData string `json:"annotationData,omitempty"`
	// Gets the validation errors from the reporting task. These validation errors represent the problems with the reporting task that must be resolved before it can be scheduled to run.
	ValidationErrors []string `json:"validationErrors,omitempty"`
	// Indicates whether the Processor is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the Processor is valid)
	ValidationStatus string `json:"validationStatus,omitempty"`
	// The number of active threads for the reporting task.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
	// Whether the underlying extension is missing.
	ExtensionMissing bool `json:"extensionMissing,omitempty"`
}
