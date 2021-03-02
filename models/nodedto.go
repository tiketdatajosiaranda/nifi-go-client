/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// NodeDto struct for NodeDto
type NodeDto struct {
	// The id of the node.
	NodeId string `json:"nodeId,omitempty"`
	// The node's host/ip address.
	Address string `json:"address,omitempty"`
	// The port the node is listening for API requests.
	ApiPort int32 `json:"apiPort,omitempty"`
	// The node's status.
	Status string `json:"status,omitempty"`
	// the time of the nodes's last heartbeat.
	Heartbeat string `json:"heartbeat,omitempty"`
	// The time of the node's last connection request.
	ConnectionRequested string `json:"connectionRequested,omitempty"`
	// The roles of this node.
	Roles []string `json:"roles,omitempty"`
	// The active threads for the NiFi on the node.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
	// The queue the NiFi on the node.
	Queued string `json:"queued,omitempty"`
	// The node's events.
	Events []NodeEventDto `json:"events,omitempty"`
	// The time at which this Node was last refreshed.
	NodeStartTime string `json:"nodeStartTime,omitempty"`
}
