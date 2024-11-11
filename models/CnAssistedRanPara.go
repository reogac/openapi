package models

type CnAssistedRanPara struct {
	TrafficProfile             string                      `json:"trafficProfile,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	StationaryIndication       string                      `json:"stationaryIndication,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ScheduledCommunicationType string                      `json:"scheduledCommunicationType,omitempty"`
}
