/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExpectedUeBehaviourData struct {
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
}
