package models

type ThresholdLevel struct {
	CongLevel         *int     `json:"congLevel,omitempty"`
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
}
