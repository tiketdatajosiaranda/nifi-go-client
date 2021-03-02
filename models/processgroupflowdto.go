/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProcessGroupFlowDto struct for ProcessGroupFlowDto
type ProcessGroupFlowDto struct {
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri string `json:"uri,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId    string                          `json:"parentGroupId,omitempty"`
	ParameterContext ParameterContextReferenceEntity `json:"parameterContext,omitempty"`
	Breadcrumb       FlowBreadcrumbEntity            `json:"breadcrumb,omitempty"`
	Flow             FlowDto                         `json:"flow,omitempty"`
	// The time the flow for the process group was last refreshed.
	LastRefreshed string `json:"lastRefreshed,omitempty"`
}
