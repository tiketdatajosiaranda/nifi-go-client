/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// FlowDto struct for FlowDto
type FlowDto struct {
	// The process groups in this flow.
	ProcessGroups []ProcessGroupEntity `json:"processGroups,omitempty"`
	// The remote process groups in this flow.
	RemoteProcessGroups []RemoteProcessGroupEntity `json:"remoteProcessGroups,omitempty"`
	// The processors in this flow.
	Processors []ProcessorEntity `json:"processors,omitempty"`
	// The input ports in this flow.
	InputPorts []PortEntity `json:"inputPorts,omitempty"`
	// The output ports in this flow.
	OutputPorts []PortEntity `json:"outputPorts,omitempty"`
	// The connections in this flow.
	Connections []ConnectionEntity `json:"connections,omitempty"`
	// The labels in this flow.
	Labels []LabelEntity `json:"labels,omitempty"`
	// The funnels in this flow.
	Funnels []FunnelEntity `json:"funnels,omitempty"`
}
