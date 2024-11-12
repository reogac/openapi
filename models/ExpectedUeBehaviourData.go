package models

type ExpectedUeBehaviourData struct {
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
}
