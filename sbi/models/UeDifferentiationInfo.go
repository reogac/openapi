/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeDifferentiationInfo struct {
	ValidityTime     string                         `json:"validityTime,omitempty"`
	PeriodicComInd   PeriodicCommunicationIndicator `json:"periodicComInd,omitempty"`
	PeriodicTime     *int                           `json:"periodicTime,omitempty"`
	ScheduledComTime *ScheduledCommunicationTime    `json:"scheduledComTime,omitempty"`
	StationaryInd    StationaryIndication           `json:"stationaryInd,omitempty"`
	TrafficProfile   TrafficProfile                 `json:"trafficProfile,omitempty"`
	BatteryInd       *BatteryIndication             `json:"batteryInd,omitempty"`
}
