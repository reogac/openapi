package models
type LteV2xAuth struct {
	 VehicleUeAuth	UeAuth	`json:"vehicleUeAuth,omitempty"`
	 PedestrianUeAuth	UeAuth	`json:"pedestrianUeAuth,omitempty"`
}
