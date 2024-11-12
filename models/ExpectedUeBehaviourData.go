package models

type ExpectedUeBehaviourData struct {
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
}
