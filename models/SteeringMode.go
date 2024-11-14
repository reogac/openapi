/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SteeringMode struct {
	ThresValue     *ThresholdValue    `json:"thresValue,omitempty"`
	SteerModeInd   SteerModeIndicator `json:"steerModeInd,omitempty"`
	SteerModeValue SteerModeValue     `json:"steerModeValue"`
	Active         AccessType         `json:"active,omitempty"`
	Standby        AccessType         `json:"standby,omitempty"`
	ThreeGLoad     *int               `json:"3gLoad,omitempty"`
	PrioAcc        AccessType         `json:"prioAcc,omitempty"`
}
