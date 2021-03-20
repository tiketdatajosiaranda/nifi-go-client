/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// VersionedProcessGroup struct for VersionedProcessGroup
type VersionedProcessGroup struct {
	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`
	// The component's name
	Name string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments string   `json:"comments,omitempty"`
	Position Position `json:"position,omitempty"`
	// The child Process Groups
	ProcessGroups []VersionedProcessGroup `json:"processGroups,omitempty"`
	// The Remote Process Groups
	RemoteProcessGroups []VersionedRemoteProcessGroup `json:"remoteProcessGroups,omitempty"`
	// The Processors
	Processors []VersionedProcessor `json:"processors,omitempty"`
	// The Input Ports
	InputPorts []VersionedPort `json:"inputPorts,omitempty"`
	// The Output Ports
	OutputPorts []VersionedPort `json:"outputPorts,omitempty"`
	// The Connections
	Connections []VersionedConnection `json:"connections,omitempty"`
	// The Labels
	Labels []VersionedLabel `json:"labels,omitempty"`
	// The Funnels
	Funnels []VersionedFunnel `json:"funnels,omitempty"`
	// The Controller Services
	ControllerServices       []VersionedControllerService `json:"controllerServices,omitempty"`
	VersionedFlowCoordinates *VersionedFlowCoordinates    `json:"versionedFlowCoordinates,omitempty"`
	// The Variables in the Variable Registry for this Process Group (not including any ancestor or descendant Process Groups)
	Variables map[string]string `json:"variables,omitempty"`
	// The name of the parameter context used by this process group
	ParameterContextName string `json:"parameterContextName,omitempty"`
	ComponentType        string `json:"componentType,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`
}
