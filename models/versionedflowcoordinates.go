/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// VersionedFlowCoordinates struct for VersionedFlowCoordinates
type VersionedFlowCoordinates struct {
	// The URL of the Flow Registry that contains the flow
	RegistryUrl string `json:"registryUrl,omitempty"`
	// The UUID of the bucket that the flow resides in
	BucketId string `json:"bucketId,omitempty"`
	// The UUID of the flow
	FlowId string `json:"flowId,omitempty"`
	// The version of the flow
	Version int32 `json:"version,omitempty"`
	// Whether or not these coordinates point to the latest version of the flow
	Latest bool `json:"latest,omitempty"`
}
