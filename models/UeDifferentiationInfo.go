package models
type UeDifferentiationInfo struct {
	 TrafficProfile	TrafficProfile	`json:"trafficProfile,omitempty"`
	 BatteryInd	*BatteryIndication	`json:"batteryInd,omitempty"`
	 ValidityTime	string	`json:"validityTime,omitempty"`
	 PeriodicComInd	PeriodicCommunicationIndicator	`json:"periodicComInd,omitempty"`
	 PeriodicTime	*int	`json:"periodicTime,omitempty"`
	 ScheduledComTime	*ScheduledCommunicationTime	`json:"scheduledComTime,omitempty"`
	 StationaryInd	StationaryIndication	`json:"stationaryInd,omitempty"`
}
