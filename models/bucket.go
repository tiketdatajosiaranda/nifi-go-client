/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// Bucket struct for Bucket
type Bucket struct {
	Link JaxbLink `json:"link,omitempty"`
	// An ID to uniquely identify this object.
	Identifier string `json:"identifier,omitempty"`
	// The name of the bucket.
	Name string `json:"name"`
	// The timestamp of when the bucket was first created. This is set by the server at creation time.
	CreatedTimestamp int64 `json:"createdTimestamp,omitempty"`
	// A description of the bucket.
	Description string `json:"description,omitempty"`
	// Indicates if this bucket allows the same version of an extension bundle to be redeployed and thus overwrite the existing artifact. By default this is false.
	AllowBundleRedeploy bool `json:"allowBundleRedeploy,omitempty"`
	// Indicates if this bucket allows read access to unauthenticated anonymous users
	AllowPublicRead bool        `json:"allowPublicRead,omitempty"`
	Permissions     Permissions `json:"permissions,omitempty"`
}
