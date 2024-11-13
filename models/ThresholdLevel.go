package models

type ThresholdLevel struct {
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
	CongLevel         *int     `json:"congLevel,omitempty"`
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
}
