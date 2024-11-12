package models

type ExpectedUeBehaviourData struct {
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	TrafficProfile             string                      `json:"trafficProfile,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationType string                      `json:"scheduledCommunicationType,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	StationaryIndication       string                      `json:"stationaryIndication,omitempty"`
}
