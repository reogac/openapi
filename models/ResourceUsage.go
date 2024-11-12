package models

type ResourceUsage struct {
	StorageUsage *int `json:"storageUsage,omitempty"`
	CpuUsage     *int `json:"cpuUsage,omitempty"`
	MemoryUsage  *int `json:"memoryUsage,omitempty"`
}
