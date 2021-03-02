/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// VersionedFlowSnapshotMetadata struct for VersionedFlowSnapshotMetadata
type VersionedFlowSnapshotMetadata struct {
	Link JaxbLink `json:"link,omitempty"`
	// The identifier of the bucket this snapshot belongs to.
	BucketIdentifier string `json:"bucketIdentifier"`
	// The identifier of the flow this snapshot belongs to.
	FlowIdentifier string `json:"flowIdentifier"`
	// The version of this snapshot of the flow.
	Version int32 `json:"version"`
	// The timestamp when the flow was saved, as milliseconds since epoch.
	Timestamp int64 `json:"timestamp,omitempty"`
	// The user that created this snapshot of the flow.
	Author string `json:"author,omitempty"`
	// The comments provided by the user when creating the snapshot.
	Comments string `json:"comments,omitempty"`
}