package models

type LteV2xAuth struct {
	PedestrianUeAuth string `json:"pedestrianUeAuth,omitempty"`
	VehicleUeAuth    string `json:"vehicleUeAuth,omitempty"`
}
