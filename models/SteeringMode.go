package models
type SteeringMode struct {
	 Active	AccessType	`json:"active,omitempty"`
	 Standby	AccessType	`json:"standby,omitempty"`
	 ThreeGLoad	*int	`json:"3gLoad,omitempty"`
	 PrioAcc	AccessType	`json:"prioAcc,omitempty"`
	 ThresValue	*ThresholdValue	`json:"thresValue,omitempty"`
	 SteerModeInd	SteerModeIndicator	`json:"steerModeInd,omitempty"`
	 SteerModeValue	SteerModeValue	`json:"steerModeValue"`
}
