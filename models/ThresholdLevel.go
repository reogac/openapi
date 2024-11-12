package models

type ThresholdLevel struct {
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
	CongLevel         *int     `json:"congLevel,omitempty"`
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
}
