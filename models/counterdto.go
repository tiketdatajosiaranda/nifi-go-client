/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// CounterDto struct for CounterDto
type CounterDto struct {
	// The id of the counter.
	Id string `json:"id,omitempty"`
	// The context of the counter.
	Context string `json:"context,omitempty"`
	// The name of the counter.
	Name string `json:"name,omitempty"`
	// The value count.
	ValueCount int64 `json:"valueCount,omitempty"`
	// The value of the counter.
	Value string `json:"value,omitempty"`
}
