/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// BulletinDto struct for BulletinDto
type BulletinDto struct {
	// The id of the bulletin.
	Id int64 `json:"id,omitempty"`
	// If clustered, the address of the node from which the bulletin originated.
	NodeAddress string `json:"nodeAddress,omitempty"`
	// The category of this bulletin.
	Category string `json:"category,omitempty"`
	// The group id of the source component.
	GroupId string `json:"groupId,omitempty"`
	// The id of the source component.
	SourceId string `json:"sourceId,omitempty"`
	// The name of the source component.
	SourceName string `json:"sourceName,omitempty"`
	// The level of the bulletin.
	Level string `json:"level,omitempty"`
	// The bulletin message.
	Message string `json:"message,omitempty"`
	// When this bulletin was generated.
	Timestamp string `json:"timestamp,omitempty"`
}
