/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ComponentSearchResultDto struct for ComponentSearchResultDto
type ComponentSearchResultDto struct {
	// The id of the component that matched the search.
	Id string `json:"id,omitempty"`
	// The group id of the component that matched the search.
	GroupId        string               `json:"groupId,omitempty"`
	ParentGroup    SearchResultGroupDto `json:"parentGroup,omitempty"`
	VersionedGroup SearchResultGroupDto `json:"versionedGroup,omitempty"`
	// The name of the component that matched the search.
	Name string `json:"name,omitempty"`
	// What matched the search from the component.
	Matches []string `json:"matches,omitempty"`
}
