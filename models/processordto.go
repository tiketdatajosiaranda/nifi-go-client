/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProcessorDto struct for ProcessorDto
type ProcessorDto struct {
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string      `json:"parentGroupId,omitempty"`
	Position      PositionDto `json:"position,omitempty"`
	// The name of the processor.
	Name string `json:"name,omitempty"`
	// The type of the processor.
	Type   string    `json:"type,omitempty"`
	Bundle BundleDto `json:"bundle,omitempty"`
	// The state of the processor
	State string `json:"state,omitempty"`
	// Styles for the processor (background-color : #eee).
	Style map[string]string `json:"style,omitempty"`
	// The available relationships that the processor currently supports.
	Relationships []RelationshipDto `json:"relationships,omitempty"`
	// The description of the processor.
	Description string `json:"description,omitempty"`
	// Whether the processor supports parallel processing.
	SupportsParallelProcessing bool `json:"supportsParallelProcessing,omitempty"`
	// Whether the processor supports event driven scheduling.
	SupportsEventDriven bool `json:"supportsEventDriven,omitempty"`
	// Whether the processor supports batching. This makes the run duration settings available.
	SupportsBatching bool `json:"supportsBatching,omitempty"`
	// Whether the processor persists state.
	PersistsState bool `json:"persistsState,omitempty"`
	// Whether the processor requires elevated privileges.
	Restricted bool `json:"restricted,omitempty"`
	// Whether the processor has been deprecated.
	Deprecated bool `json:"deprecated,omitempty"`
	// Indicates if the execution node of a processor is restricted to run only on the primary node
	ExecutionNodeRestricted bool `json:"executionNodeRestricted,omitempty"`
	// Whether the processor has multiple versions available.
	MultipleVersionsAvailable bool `json:"multipleVersionsAvailable,omitempty"`
	// The input requirement for this processor.
	InputRequirement string             `json:"inputRequirement,omitempty"`
	Config           ProcessorConfigDto `json:"config,omitempty"`
	// The validation errors for the processor. These validation errors represent the problems with the processor that must be resolved before it can be started.
	ValidationErrors []string `json:"validationErrors,omitempty"`
	// Indicates whether the Processor is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the Processor is valid)
	ValidationStatus string `json:"validationStatus,omitempty"`
	// Whether the underlying extension is missing.
	ExtensionMissing bool `json:"extensionMissing,omitempty"`
}
