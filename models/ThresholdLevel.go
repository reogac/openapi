package models

type ThresholdLevel struct {
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
	CongLevel         *int     `json:"congLevel,omitempty"`
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
}
