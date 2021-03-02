/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// SystemDiagnosticsSnapshotDto struct for SystemDiagnosticsSnapshotDto
type SystemDiagnosticsSnapshotDto struct {
	// Total size of non heap.
	TotalNonHeap string `json:"totalNonHeap,omitempty"`
	// Total number of bytes allocated to the JVM not used for heap
	TotalNonHeapBytes int64 `json:"totalNonHeapBytes,omitempty"`
	// Amount of use non heap.
	UsedNonHeap string `json:"usedNonHeap,omitempty"`
	// Total number of bytes used by the JVM not in the heap space
	UsedNonHeapBytes int64 `json:"usedNonHeapBytes,omitempty"`
	// Amount of free non heap.
	FreeNonHeap string `json:"freeNonHeap,omitempty"`
	// Total number of free non-heap bytes available to the JVM
	FreeNonHeapBytes int64 `json:"freeNonHeapBytes,omitempty"`
	// Maximum size of non heap.
	MaxNonHeap string `json:"maxNonHeap,omitempty"`
	// The maximum number of bytes that the JVM can use for non-heap purposes
	MaxNonHeapBytes int64 `json:"maxNonHeapBytes,omitempty"`
	// Utilization of non heap.
	NonHeapUtilization string `json:"nonHeapUtilization,omitempty"`
	// Total size of heap.
	TotalHeap string `json:"totalHeap,omitempty"`
	// The total number of bytes that are available for the JVM heap to use
	TotalHeapBytes int64 `json:"totalHeapBytes,omitempty"`
	// Amount of used heap.
	UsedHeap string `json:"usedHeap,omitempty"`
	// The number of bytes of JVM heap that are currently being used
	UsedHeapBytes int64 `json:"usedHeapBytes,omitempty"`
	// Amount of free heap.
	FreeHeap string `json:"freeHeap,omitempty"`
	// The number of bytes that are allocated to the JVM heap but not currently being used
	FreeHeapBytes int64 `json:"freeHeapBytes,omitempty"`
	// Maximum size of heap.
	MaxHeap string `json:"maxHeap,omitempty"`
	// The maximum number of bytes that can be used by the JVM
	MaxHeapBytes int64 `json:"maxHeapBytes,omitempty"`
	// Utilization of heap.
	HeapUtilization string `json:"heapUtilization,omitempty"`
	// Number of available processors if supported by the underlying system.
	AvailableProcessors int32 `json:"availableProcessors,omitempty"`
	// The processor load average if supported by the underlying system.
	ProcessorLoadAverage float64 `json:"processorLoadAverage,omitempty"`
	// Total number of threads.
	TotalThreads int32 `json:"totalThreads,omitempty"`
	// Number of daemon threads.
	DaemonThreads int32 `json:"daemonThreads,omitempty"`
	// The uptime of the Java virtual machine
	Uptime                         string          `json:"uptime,omitempty"`
	FlowFileRepositoryStorageUsage StorageUsageDto `json:"flowFileRepositoryStorageUsage,omitempty"`
	// The content repository storage usage.
	ContentRepositoryStorageUsage []StorageUsageDto `json:"contentRepositoryStorageUsage,omitempty"`
	// The provenance repository storage usage.
	ProvenanceRepositoryStorageUsage []StorageUsageDto `json:"provenanceRepositoryStorageUsage,omitempty"`
	// The garbage collection details.
	GarbageCollection []GarbageCollectionDto `json:"garbageCollection,omitempty"`
	// When the diagnostics were generated.
	StatsLastRefreshed string         `json:"statsLastRefreshed,omitempty"`
	VersionInfo        VersionInfoDto `json:"versionInfo,omitempty"`
}
