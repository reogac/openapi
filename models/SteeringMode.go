package models

type SteeringMode struct {
	SteerModeInd   SteerModeIndicator `json:"steerModeInd,omitempty"`
	SteerModeValue SteerModeValue     `json:"steerModeValue"`
	Active         AccessType         `json:"active,omitempty"`
	Standby        AccessType         `json:"standby,omitempty"`
	ThreeGLoad     *int               `json:"3gLoad,omitempty"`
	PrioAcc        AccessType         `json:"prioAcc,omitempty"`
	ThresValue     *ThresholdValue    `json:"thresValue,omitempty"`
}
