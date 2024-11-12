package models

type ThresholdLevel struct {
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
	CongLevel         *int     `json:"congLevel,omitempty"`
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
}
