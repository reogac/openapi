package models
type ResourceUsage struct {
	 CpuUsage	*int	`json:"cpuUsage,omitempty"`
	 MemoryUsage	*int	`json:"memoryUsage,omitempty"`
	 StorageUsage	*int	`json:"storageUsage,omitempty"`
}
