package models
type NrV2xAuth struct {
	 VehicleUeAuth	UeAuth	`json:"vehicleUeAuth,omitempty"`
	 PedestrianUeAuth	UeAuth	`json:"pedestrianUeAuth,omitempty"`
}
