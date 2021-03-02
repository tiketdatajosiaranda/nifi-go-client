/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// CurrentUserEntity struct for CurrentUserEntity
type CurrentUserEntity struct {
	// The user identity being serialized.
	Identity string `json:"identity,omitempty"`
	// Whether the current user is anonymous.
	Anonymous                       bool           `json:"anonymous,omitempty"`
	ProvenancePermissions           PermissionsDto `json:"provenancePermissions,omitempty"`
	CountersPermissions             PermissionsDto `json:"countersPermissions,omitempty"`
	TenantsPermissions              PermissionsDto `json:"tenantsPermissions,omitempty"`
	ControllerPermissions           PermissionsDto `json:"controllerPermissions,omitempty"`
	PoliciesPermissions             PermissionsDto `json:"policiesPermissions,omitempty"`
	SystemPermissions               PermissionsDto `json:"systemPermissions,omitempty"`
	ParameterContextPermissions     PermissionsDto `json:"parameterContextPermissions,omitempty"`
	RestrictedComponentsPermissions PermissionsDto `json:"restrictedComponentsPermissions,omitempty"`
	// Permissions for specific component restrictions.
	ComponentRestrictionPermissions []ComponentRestrictionPermissionDto `json:"componentRestrictionPermissions,omitempty"`
	// Whether the current user can version flows.
	CanVersionFlows bool `json:"canVersionFlows,omitempty"`
}
