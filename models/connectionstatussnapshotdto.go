/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ConnectionStatusSnapshotDto struct for ConnectionStatusSnapshotDto
type ConnectionStatusSnapshotDto struct {
	// The id of the connection.
	Id string `json:"id,omitempty"`
	// The id of the process group the connection belongs to.
	GroupId string `json:"groupId,omitempty"`
	// The name of the connection.
	Name string `json:"name,omitempty"`
	// The id of the source of the connection.
	SourceId string `json:"sourceId,omitempty"`
	// The name of the source of the connection.
	SourceName string `json:"sourceName,omitempty"`
	// The id of the destination of the connection.
	DestinationId string `json:"destinationId,omitempty"`
	// The name of the destination of the connection.
	DestinationName string                                 `json:"destinationName,omitempty"`
	Predictions     *ConnectionStatusPredictionsSnapshotDto `json:"predictions,omitempty"`
	// The number of FlowFiles that have come into the connection in the last 5 minutes.
	FlowFilesIn int32 `json:"flowFilesIn,omitempty"`
	// The size of the FlowFiles that have come into the connection in the last 5 minutes.
	BytesIn int64 `json:"bytesIn,omitempty"`
	// The input count/size for the connection in the last 5 minutes, pretty printed.
	Input string `json:"input,omitempty"`
	// The number of FlowFiles that have left the connection in the last 5 minutes.
	FlowFilesOut int32 `json:"flowFilesOut,omitempty"`
	// The number of bytes that have left the connection in the last 5 minutes.
	BytesOut int64 `json:"bytesOut,omitempty"`
	// The output count/sie for the connection in the last 5 minutes, pretty printed.
	Output string `json:"output,omitempty"`
	// The number of FlowFiles that are currently queued in the connection.
	FlowFilesQueued int32 `json:"flowFilesQueued,omitempty"`
	// The size of the FlowFiles that are currently queued in the connection.
	BytesQueued int64 `json:"bytesQueued,omitempty"`
	// The total count and size of queued flowfiles formatted.
	Queued string `json:"queued,omitempty"`
	// The total size of flowfiles that are queued formatted.
	QueuedSize string `json:"queuedSize,omitempty"`
	// The number of flowfiles that are queued, pretty printed.
	QueuedCount string `json:"queuedCount,omitempty"`
	// Connection percent use regarding queued flow files count and backpressure threshold if configured.
	PercentUseCount int32 `json:"percentUseCount,omitempty"`
	// Connection percent use regarding queued flow files size and backpressure threshold if configured.
	PercentUseBytes int32 `json:"percentUseBytes,omitempty"`
}
