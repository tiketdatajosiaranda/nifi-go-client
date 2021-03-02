/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ConnectionStatusPredictionsSnapshotDto struct for ConnectionStatusPredictionsSnapshotDto
type ConnectionStatusPredictionsSnapshotDto struct {
	// The predicted number of milliseconds before the connection will have backpressure applied, based on the queued count.
	PredictedMillisUntilCountBackpressure int64 `json:"predictedMillisUntilCountBackpressure,omitempty"`
	// The predicted number of milliseconds before the connection will have backpressure applied, based on the total number of bytes in the queue.
	PredictedMillisUntilBytesBackpressure int64 `json:"predictedMillisUntilBytesBackpressure,omitempty"`
	// The configured interval (in seconds) for predicting connection queue count and size (and percent usage).
	PredictionIntervalSeconds int32 `json:"predictionIntervalSeconds,omitempty"`
	// The predicted number of queued objects at the next configured interval.
	PredictedCountAtNextInterval int32 `json:"predictedCountAtNextInterval,omitempty"`
	// The predicted total number of bytes in the queue at the next configured interval.
	PredictedBytesAtNextInterval int64 `json:"predictedBytesAtNextInterval,omitempty"`
	// Predicted connection percent use regarding queued flow files count and backpressure threshold if configured.
	PredictedPercentCount int32 `json:"predictedPercentCount,omitempty"`
	// Predicted connection percent use regarding queued flow files size and backpressure threshold if configured.
	PredictedPercentBytes int32 `json:"predictedPercentBytes,omitempty"`
}
