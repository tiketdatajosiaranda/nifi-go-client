/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// RemoteProcessGroupContentsDto struct for RemoteProcessGroupContentsDto
type RemoteProcessGroupContentsDto struct {
	// The input ports to which data can be sent.
	InputPorts []RemoteProcessGroupPortDto `json:"inputPorts,omitempty"`
	// The output ports from which data can be retrieved.
	OutputPorts []RemoteProcessGroupPortDto `json:"outputPorts,omitempty"`
}
