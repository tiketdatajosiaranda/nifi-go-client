/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// AccessPolicyDto struct for AccessPolicyDto
type AccessPolicyDto struct {
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string      `json:"parentGroupId,omitempty"`
	Position      PositionDto `json:"position,omitempty"`
	// The resource for this access policy.
	Resource string `json:"resource,omitempty"`
	// The action associated with this access policy.
	Action             string                   `json:"action,omitempty"`
	ComponentReference ComponentReferenceEntity `json:"componentReference,omitempty"`
	// Whether this policy is configurable.
	Configurable bool `json:"configurable,omitempty"`
	// The set of user IDs associated with this access policy.
	Users []TenantEntity `json:"users,omitempty"`
	// The set of user group IDs associated with this access policy.
	UserGroups []TenantEntity `json:"userGroups,omitempty"`
}
