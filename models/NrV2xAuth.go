package models

type NrV2xAuth struct {
	PedestrianUeAuth UeAuth `json:"pedestrianUeAuth,omitempty"`
	VehicleUeAuth    UeAuth `json:"vehicleUeAuth,omitempty"`
}
