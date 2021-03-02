/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProcessorEntity struct for ProcessorEntity
type ProcessorEntity struct {
	Revision RevisionDto `json:"revision,omitempty"`
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri         string         `json:"uri,omitempty"`
	Position    PositionDto    `json:"position,omitempty"`
	Permissions PermissionsDto `json:"permissions,omitempty"`
	// The bulletins for this component.
	Bulletins []BulletinEntity `json:"bulletins,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool         `json:"disconnectedNodeAcknowledged,omitempty"`
	Component                    ProcessorDto `json:"component,omitempty"`
	// The input requirement for this processor.
	InputRequirement   string             `json:"inputRequirement,omitempty"`
	Status             ProcessorStatusDto `json:"status,omitempty"`
	OperatePermissions PermissionsDto     `json:"operatePermissions,omitempty"`
}
