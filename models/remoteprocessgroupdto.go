/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// RemoteProcessGroupDto struct for RemoteProcessGroupDto
type RemoteProcessGroupDto struct {
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string      `json:"parentGroupId,omitempty"`
	Position      PositionDto `json:"position,omitempty"`
	// The target URI of the remote process group. If target uri is not set, but uris are set, then returns the first url in the urls. If neither target uri nor uris are set, then returns null.
	TargetUri string `json:"targetUri,omitempty"`
	// The target URI of the remote process group. If target uris is not set but target uri is set, then returns a collection containing the single target uri. If neither target uris nor uris are set, then returns null.
	TargetUris string `json:"targetUris,omitempty"`
	// Whether the target is running securely.
	TargetSecure bool `json:"targetSecure,omitempty"`
	// The name of the remote process group.
	Name string `json:"name,omitempty"`
	// The comments for the remote process group.
	Comments string `json:"comments,omitempty"`
	// The time period used for the timeout when communicating with the target.
	CommunicationsTimeout string `json:"communicationsTimeout,omitempty"`
	// When yielding, this amount of time must elapse before the remote process group is scheduled again.
	YieldDuration     string `json:"yieldDuration,omitempty"`
	TransportProtocol string `json:"transportProtocol,omitempty"`
	// The local network interface to send/receive data. If not specified, any local address is used. If clustered, all nodes must have an interface with this identifier.
	LocalNetworkInterface string `json:"localNetworkInterface,omitempty"`
	ProxyHost             string `json:"proxyHost,omitempty"`
	ProxyPort             int32  `json:"proxyPort,omitempty"`
	ProxyUser             string `json:"proxyUser,omitempty"`
	ProxyPassword         string `json:"proxyPassword,omitempty"`
	// Any remote authorization issues for the remote process group.
	AuthorizationIssues []string `json:"authorizationIssues,omitempty"`
	// The validation errors for the remote process group. These validation errors represent the problems with the remote process group that must be resolved before it can transmit.
	ValidationErrors []string `json:"validationErrors,omitempty"`
	// Whether the remote process group is actively transmitting.
	Transmitting bool `json:"transmitting,omitempty"`
	// The number of remote input ports currently available on the target.
	InputPortCount int32 `json:"inputPortCount,omitempty"`
	// The number of remote output ports currently available on the target.
	OutputPortCount int32 `json:"outputPortCount,omitempty"`
	// The number of active remote input ports.
	ActiveRemoteInputPortCount int32 `json:"activeRemoteInputPortCount,omitempty"`
	// The number of inactive remote input ports.
	InactiveRemoteInputPortCount int32 `json:"inactiveRemoteInputPortCount,omitempty"`
	// The number of active remote output ports.
	ActiveRemoteOutputPortCount int32 `json:"activeRemoteOutputPortCount,omitempty"`
	// The number of inactive remote output ports.
	InactiveRemoteOutputPortCount int32 `json:"inactiveRemoteOutputPortCount,omitempty"`
	// The timestamp when this remote process group was last refreshed.
	FlowRefreshed string                        `json:"flowRefreshed,omitempty"`
	Contents      RemoteProcessGroupContentsDto `json:"contents,omitempty"`
}
