package models

type ExpectedUeBehaviourData struct {
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
}
