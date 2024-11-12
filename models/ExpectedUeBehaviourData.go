package models

type ExpectedUeBehaviourData struct {
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
}
