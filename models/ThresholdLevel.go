package models

type ThresholdLevel struct {
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
	CongLevel         *int     `json:"congLevel,omitempty"`
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
}
