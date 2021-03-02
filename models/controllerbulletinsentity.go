/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ControllerBulletinsEntity struct for ControllerBulletinsEntity
type ControllerBulletinsEntity struct {
	// System level bulletins to be reported to the user.
	Bulletins []BulletinEntity `json:"bulletins,omitempty"`
	// Controller service bulletins to be reported to the user.
	ControllerServiceBulletins []BulletinEntity `json:"controllerServiceBulletins,omitempty"`
	// Reporting task bulletins to be reported to the user.
	ReportingTaskBulletins []BulletinEntity `json:"reportingTaskBulletins,omitempty"`
}
