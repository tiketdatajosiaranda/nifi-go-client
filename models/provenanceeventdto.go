/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProvenanceEventDto struct for ProvenanceEventDto
type ProvenanceEventDto struct {
	// The event uuid.
	Id string `json:"id,omitempty"`
	// The event id. This is a one up number thats unique per node.
	EventId int64 `json:"eventId,omitempty"`
	// The timestamp of the event.
	EventTime string `json:"eventTime,omitempty"`
	// The event duration in milliseconds.
	EventDuration int64 `json:"eventDuration,omitempty"`
	// The duration since the lineage began, in milliseconds.
	LineageDuration int64 `json:"lineageDuration,omitempty"`
	// The type of the event.
	EventType string `json:"eventType,omitempty"`
	// The uuid of the flowfile for the event.
	FlowFileUuid string `json:"flowFileUuid,omitempty"`
	// The size of the flowfile for the event.
	FileSize string `json:"fileSize,omitempty"`
	// The size of the flowfile in bytes for the event.
	FileSizeBytes int64 `json:"fileSizeBytes,omitempty"`
	// The identifier for the node where the event originated.
	ClusterNodeId string `json:"clusterNodeId,omitempty"`
	// The label for the node where the event originated.
	ClusterNodeAddress string `json:"clusterNodeAddress,omitempty"`
	// The id of the group that the component resides in. If the component is no longer in the flow, the group id will not be set.
	GroupId string `json:"groupId,omitempty"`
	// The id of the component that generated the event.
	ComponentId string `json:"componentId,omitempty"`
	// The type of the component that generated the event.
	ComponentType string `json:"componentType,omitempty"`
	// The name of the component that generated the event.
	ComponentName string `json:"componentName,omitempty"`
	// The source system flowfile id.
	SourceSystemFlowFileId string `json:"sourceSystemFlowFileId,omitempty"`
	// The alternate identifier uri for the fileflow for the event.
	AlternateIdentifierUri string `json:"alternateIdentifierUri,omitempty"`
	// The attributes of the flowfile for the event.
	Attributes []AttributeDto `json:"attributes,omitempty"`
	// The parent uuids for the event.
	ParentUuids []string `json:"parentUuids,omitempty"`
	// The child uuids for the event.
	ChildUuids []string `json:"childUuids,omitempty"`
	// The source/destination system uri if the event was a RECEIVE/SEND.
	TransitUri string `json:"transitUri,omitempty"`
	// The relationship to which the flowfile was routed if the event is of type ROUTE.
	Relationship string `json:"relationship,omitempty"`
	// The event details.
	Details string `json:"details,omitempty"`
	// Whether the input and output content claim is the same.
	ContentEqual bool `json:"contentEqual,omitempty"`
	// Whether the input content is still available.
	InputContentAvailable bool `json:"inputContentAvailable,omitempty"`
	// The section in which the input content claim lives.
	InputContentClaimSection string `json:"inputContentClaimSection,omitempty"`
	// The container in which the input content claim lives.
	InputContentClaimContainer string `json:"inputContentClaimContainer,omitempty"`
	// The identifier of the input content claim.
	InputContentClaimIdentifier string `json:"inputContentClaimIdentifier,omitempty"`
	// The offset into the input content claim where the flowfiles content begins.
	InputContentClaimOffset int64 `json:"inputContentClaimOffset,omitempty"`
	// The file size of the input content claim formatted.
	InputContentClaimFileSize string `json:"inputContentClaimFileSize,omitempty"`
	// The file size of the intput content claim in bytes.
	InputContentClaimFileSizeBytes int64 `json:"inputContentClaimFileSizeBytes,omitempty"`
	// Whether the output content is still available.
	OutputContentAvailable bool `json:"outputContentAvailable,omitempty"`
	// The section in which the output content claim lives.
	OutputContentClaimSection string `json:"outputContentClaimSection,omitempty"`
	// The container in which the output content claim lives.
	OutputContentClaimContainer string `json:"outputContentClaimContainer,omitempty"`
	// The identifier of the output content claim.
	OutputContentClaimIdentifier string `json:"outputContentClaimIdentifier,omitempty"`
	// The offset into the output content claim where the flowfiles content begins.
	OutputContentClaimOffset int64 `json:"outputContentClaimOffset,omitempty"`
	// The file size of the output content claim formatted.
	OutputContentClaimFileSize string `json:"outputContentClaimFileSize,omitempty"`
	// The file size of the output content claim in bytes.
	OutputContentClaimFileSizeBytes int64 `json:"outputContentClaimFileSizeBytes,omitempty"`
	// Whether or not replay is available.
	ReplayAvailable bool `json:"replayAvailable,omitempty"`
	// Explanation as to why replay is unavailable.
	ReplayExplanation string `json:"replayExplanation,omitempty"`
	// The identifier of the queue/connection from which the flowfile was pulled to genereate this event. May be null if the queue/connection is unknown or the flowfile was generated from this event.
	SourceConnectionIdentifier string `json:"sourceConnectionIdentifier,omitempty"`
}
