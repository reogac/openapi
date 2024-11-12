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
