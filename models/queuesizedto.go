/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// QueueSizeDto struct for QueueSizeDto
type QueueSizeDto struct {
	// The size of objects in a queue.
	ByteCount int64 `json:"byteCount,omitempty"`
	// The count of objects in a queue.
	ObjectCount int32 `json:"objectCount,omitempty"`
}
