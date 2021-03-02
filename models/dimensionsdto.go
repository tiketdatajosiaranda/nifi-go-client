/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// DimensionsDto struct for DimensionsDto
type DimensionsDto struct {
	// The width of the label in pixels when at a 1:1 scale.
	Width float64 `json:"width,omitempty"`
	// The height of the label in pixels when at a 1:1 scale.
	Height float64 `json:"height,omitempty"`
}
