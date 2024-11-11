package models

type ThresholdLevel struct {
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
	CongLevel         *int     `json:"congLevel,omitempty"`
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
}
