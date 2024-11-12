package models
type PerfData struct {
	 AvePacketDelay	*int	`json:"avePacketDelay,omitempty"`
	 MaxPacketDelay	*int	`json:"maxPacketDelay,omitempty"`
	 AvgPacketLossRate	*int	`json:"avgPacketLossRate,omitempty"`
	 AvgTrafficRate	string	`json:"avgTrafficRate,omitempty"`
	 MaxTrafficRate	string	`json:"maxTrafficRate,omitempty"`
}
