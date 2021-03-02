/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// FlowComparisonEntity struct for FlowComparisonEntity
type FlowComparisonEntity struct {
	// The list of differences for each component in the flow that is not the same between the two flows
	ComponentDifferences []ComponentDifferenceDto `json:"componentDifferences,omitempty"`
}
