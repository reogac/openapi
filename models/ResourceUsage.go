package models

type ResourceUsage struct {
	MemoryUsage  *int `json:"memoryUsage,omitempty"`
	StorageUsage *int `json:"storageUsage,omitempty"`
	CpuUsage     *int `json:"cpuUsage,omitempty"`
}
