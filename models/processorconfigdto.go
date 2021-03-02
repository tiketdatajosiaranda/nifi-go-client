/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProcessorConfigDto struct for ProcessorConfigDto
type ProcessorConfigDto struct {
	// The properties for the processor. Properties whose value is not set will only contain the property name.
	Properties map[string]string `json:"properties,omitempty"`
	// Descriptors for the processor's properties.
	Descriptors map[string]PropertyDescriptorDto `json:"descriptors,omitempty"`
	// The frequency with which to schedule the processor. The format of the value will depend on th value of schedulingStrategy.
	SchedulingPeriod string `json:"schedulingPeriod,omitempty"`
	// Indcates whether the prcessor should be scheduled to run in event or timer driven mode.
	SchedulingStrategy string `json:"schedulingStrategy,omitempty"`
	// Indicates the node where the process will execute.
	ExecutionNode string `json:"executionNode,omitempty"`
	// The amount of time that is used when the process penalizes a flowfile.
	PenaltyDuration string `json:"penaltyDuration,omitempty"`
	// The amount of time that must elapse before this processor is scheduled again after yielding.
	YieldDuration string `json:"yieldDuration,omitempty"`
	// The level at which the processor will report bulletins.
	BulletinLevel string `json:"bulletinLevel,omitempty"`
	// The run duration for the processor in milliseconds.
	RunDurationMillis int64 `json:"runDurationMillis,omitempty"`
	// The number of tasks that should be concurrently schedule for the processor. If the processor doesn't allow parallol processing then any positive input will be ignored.
	ConcurrentlySchedulableTaskCount int32 `json:"concurrentlySchedulableTaskCount,omitempty"`
	// The names of all relationships that cause a flow file to be terminated if the relationship is not connected elsewhere. This property differs from the 'isAutoTerminate' property of the RelationshipDTO in that the RelationshipDTO is meant to depict the current configuration, whereas this property can be set in a DTO when updating a Processor in order to change which Relationships should be auto-terminated.
	AutoTerminatedRelationships []string `json:"autoTerminatedRelationships,omitempty"`
	// The comments for the processor.
	Comments string `json:"comments,omitempty"`
	// The URL for the processor's custom configuration UI if applicable.
	CustomUiUrl string `json:"customUiUrl,omitempty"`
	// Whether the processor is loss tolerant.
	LossTolerant bool `json:"lossTolerant,omitempty"`
	// The annotation data for the processor used to relay configuration between a custom UI and the procesosr.
	AnnotationData string `json:"annotationData,omitempty"`
	// Maps default values for concurrent tasks for each applicable scheduling strategy.
	DefaultConcurrentTasks map[string]string `json:"defaultConcurrentTasks,omitempty"`
	// Maps default values for scheduling period for each applicable scheduling strategy.
	DefaultSchedulingPeriod map[string]string `json:"defaultSchedulingPeriod,omitempty"`
}
