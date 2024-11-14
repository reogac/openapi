package models
type TrafficInformation struct {
	 UplinkRate	string	`json:"uplinkRate,omitempty"`
	 DownlinkRate	string	`json:"downlinkRate,omitempty"`
	 UplinkVolume	*int64	`json:"uplinkVolume,omitempty"`
	 DownlinkVolume	*int64	`json:"downlinkVolume,omitempty"`
	 TotalVolume	*int64	`json:"totalVolume,omitempty"`
}
