package models

type LteV2xAuth struct {
	VehicleUeAuth    string `json:"vehicleUeAuth,omitempty"`
	PedestrianUeAuth string `json:"pedestrianUeAuth,omitempty"`
}
