/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// PeerDto struct for PeerDto
type PeerDto struct {
	// The hostname of this peer.
	Hostname string `json:"hostname,omitempty"`
	// The port number of this peer.
	Port int32 `json:"port,omitempty"`
	// Returns if this peer connection is secure.
	Secure bool `json:"secure,omitempty"`
	// The number of flowFiles this peer holds.
	FlowFileCount int32 `json:"flowFileCount,omitempty"`
}
