package models

type UeDifferentiationInfo struct {
	BatteryInd       *BatteryIndication          `json:"batteryInd,omitempty"`
	ValidityTime     string                      `json:"validityTime,omitempty"`
	PeriodicComInd   string                      `json:"periodicComInd,omitempty"`
	PeriodicTime     *int                        `json:"periodicTime,omitempty"`
	ScheduledComTime *ScheduledCommunicationTime `json:"scheduledComTime,omitempty"`
	StationaryInd    string                      `json:"stationaryInd,omitempty"`
	TrafficProfile   string                      `json:"trafficProfile,omitempty"`
}
