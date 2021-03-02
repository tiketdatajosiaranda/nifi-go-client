/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// BucketDto struct for BucketDto
type BucketDto struct {
	// The bucket identifier
	Id string `json:"id,omitempty"`
	// The bucket name
	Name string `json:"name,omitempty"`
	// The bucket description
	Description string `json:"description,omitempty"`
	// The created timestamp of this bucket
	Created int64 `json:"created,omitempty"`
}
