/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// AboutDto struct for AboutDto
type AboutDto struct {
	// The title to be used on the page and in the about dialog.
	Title string `json:"title,omitempty"`
	// The version of this NiFi.
	Version string `json:"version,omitempty"`
	// The URI for the NiFi.
	Uri string `json:"uri,omitempty"`
	// The URL for the content viewer if configured.
	ContentViewerUrl string `json:"contentViewerUrl,omitempty"`
	// The timezone of the NiFi instance.
	Timezone string `json:"timezone,omitempty"`
	// Build tag
	BuildTag string `json:"buildTag,omitempty"`
	// Build revision or commit hash
	BuildRevision string `json:"buildRevision,omitempty"`
	// Build branch
	BuildBranch string `json:"buildBranch,omitempty"`
	// Build timestamp
	BuildTimestamp string `json:"buildTimestamp,omitempty"`
}
