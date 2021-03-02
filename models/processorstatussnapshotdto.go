/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// ProcessorStatusSnapshotDto struct for ProcessorStatusSnapshotDto
type ProcessorStatusSnapshotDto struct {
	// The id of the processor.
	Id string `json:"id,omitempty"`
	// The id of the parent process group to which the processor belongs.
	GroupId string `json:"groupId,omitempty"`
	// The name of the prcessor.
	Name string `json:"name,omitempty"`
	// The type of the processor.
	Type string `json:"type,omitempty"`
	// The state of the processor.
	RunStatus string `json:"runStatus,omitempty"`
	// Indicates the node where the process will execute.
	ExecutionNode string `json:"executionNode,omitempty"`
	// The number of bytes read by this Processor in the last 5 mintues
	BytesRead int64 `json:"bytesRead,omitempty"`
	// The number of bytes written by this Processor in the last 5 minutes
	BytesWritten int64 `json:"bytesWritten,omitempty"`
	// The number of bytes read in the last 5 minutes.
	Read string `json:"read,omitempty"`
	// The number of bytes written in the last 5 minutes.
	Written string `json:"written,omitempty"`
	// The number of FlowFiles that have been accepted in the last 5 minutes
	FlowFilesIn int32 `json:"flowFilesIn,omitempty"`
	// The size of the FlowFiles that have been accepted in the last 5 minutes
	BytesIn int64 `json:"bytesIn,omitempty"`
	// The count/size of flowfiles that have been accepted in the last 5 minutes.
	Input string `json:"input,omitempty"`
	// The number of FlowFiles transferred to a Connection in the last 5 minutes
	FlowFilesOut int32 `json:"flowFilesOut,omitempty"`
	// The size of the FlowFiles transferred to a Connection in the last 5 minutes
	BytesOut int64 `json:"bytesOut,omitempty"`
	// The count/size of flowfiles that have been processed in the last 5 minutes.
	Output string `json:"output,omitempty"`
	// The number of times this Processor has run in the last 5 minutes
	TaskCount int32 `json:"taskCount,omitempty"`
	// The number of nanoseconds that this Processor has spent running in the last 5 minutes
	TasksDurationNanos int64 `json:"tasksDurationNanos,omitempty"`
	// The total number of task this connectable has completed over the last 5 minutes.
	Tasks string `json:"tasks,omitempty"`
	// The total duration of all tasks for this connectable over the last 5 minutes.
	TasksDuration string `json:"tasksDuration,omitempty"`
	// The number of threads currently executing in the processor.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
	// The number of threads currently terminated for the processor.
	TerminatedThreadCount int32 `json:"terminatedThreadCount,omitempty"`
}
