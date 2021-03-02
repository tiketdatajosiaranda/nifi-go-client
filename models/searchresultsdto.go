/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// SearchResultsDto struct for SearchResultsDto
type SearchResultsDto struct {
	// The processors that matched the search.
	ProcessorResults []ComponentSearchResultDto `json:"processorResults,omitempty"`
	// The connections that matched the search.
	ConnectionResults []ComponentSearchResultDto `json:"connectionResults,omitempty"`
	// The process groups that matched the search.
	ProcessGroupResults []ComponentSearchResultDto `json:"processGroupResults,omitempty"`
	// The input ports that matched the search.
	InputPortResults []ComponentSearchResultDto `json:"inputPortResults,omitempty"`
	// The output ports that matched the search.
	OutputPortResults []ComponentSearchResultDto `json:"outputPortResults,omitempty"`
	// The remote process groups that matched the search.
	RemoteProcessGroupResults []ComponentSearchResultDto `json:"remoteProcessGroupResults,omitempty"`
	// The funnels that matched the search.
	FunnelResults []ComponentSearchResultDto `json:"funnelResults,omitempty"`
	// The labels that matched the search.
	LabelResults []ComponentSearchResultDto `json:"labelResults,omitempty"`
	// The controller service nodes that matched the search
	ControllerServiceNodeResults []ComponentSearchResultDto `json:"controllerServiceNodeResults,omitempty"`
	// The parameter contexts that matched the search.
	ParameterContextResults []ComponentSearchResultDto `json:"parameterContextResults,omitempty"`
	// The parameters that matched the search.
	ParameterResults []ComponentSearchResultDto `json:"parameterResults,omitempty"`
}
