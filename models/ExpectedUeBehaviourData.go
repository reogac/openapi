package models

type ExpectedUeBehaviourData struct {
	TrafficProfile             string                      `json:"trafficProfile,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	ScheduledCommunicationType string                      `json:"scheduledCommunicationType,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	StationaryIndication       string                      `json:"stationaryIndication,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
}
