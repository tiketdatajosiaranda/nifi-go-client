/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// FlowSnippetDto struct for FlowSnippetDto
type FlowSnippetDto struct {
	// The process groups in this flow snippet.
	ProcessGroups []ProcessGroupDto `json:"processGroups,omitempty"`
	// The remote process groups in this flow snippet.
	RemoteProcessGroups []RemoteProcessGroupDto `json:"remoteProcessGroups,omitempty"`
	// The processors in this flow snippet.
	Processors []ProcessorDto `json:"processors,omitempty"`
	// The input ports in this flow snippet.
	InputPorts []PortDto `json:"inputPorts,omitempty"`
	// The output ports in this flow snippet.
	OutputPorts []PortDto `json:"outputPorts,omitempty"`
	// The connections in this flow snippet.
	Connections []ConnectionDto `json:"connections,omitempty"`
	// The labels in this flow snippet.
	Labels []LabelDto `json:"labels,omitempty"`
	// The funnels in this flow snippet.
	Funnels []FunnelDto `json:"funnels,omitempty"`
	// The controller services in this flow snippet.
	ControllerServices []ControllerServiceDto `json:"controllerServices,omitempty"`
}
