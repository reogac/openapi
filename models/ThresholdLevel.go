package models
type ThresholdLevel struct {
	 AvgTrafficRate	string	`json:"avgTrafficRate,omitempty"`
	 SvcExpLevel	*float64	`json:"svcExpLevel,omitempty"`
	 NfCpuUsage	*int	`json:"nfCpuUsage,omitempty"`
	 NfStorageUsage	*int	`json:"nfStorageUsage,omitempty"`
	 NfMemoryUsage	*int	`json:"nfMemoryUsage,omitempty"`
	 MaxTrafficRate	string	`json:"maxTrafficRate,omitempty"`
	 AvgPacketDelay	*int	`json:"avgPacketDelay,omitempty"`
	 MaxPacketDelay	*int	`json:"maxPacketDelay,omitempty"`
	 AvgPacketLossRate	*int	`json:"avgPacketLossRate,omitempty"`
	 CongLevel	*int	`json:"congLevel,omitempty"`
	 NfLoadLevel	*int	`json:"nfLoadLevel,omitempty"`
}
