package models
type ExpectedUeBehaviourData struct {
	 ScheduledCommunicationTime	*ScheduledCommunicationTime	`json:"scheduledCommunicationTime,omitempty"`
	 ScheduledCommunicationType	ScheduledCommunicationType	`json:"scheduledCommunicationType,omitempty"`
	 StationaryIndication	StationaryIndication	`json:"stationaryIndication,omitempty"`
	 CommunicationDurationTime	*int	`json:"communicationDurationTime,omitempty"`
	 PeriodicTime	*int	`json:"periodicTime,omitempty"`
	 ValidityTime	string	`json:"validityTime,omitempty"`
	 ExpectedUmts	[]LocationArea	`json:"expectedUmts,omitempty"`
	 TrafficProfile	TrafficProfile	`json:"trafficProfile,omitempty"`
	 BatteryIndication	*BatteryIndication	`json:"batteryIndication,omitempty"`
}
