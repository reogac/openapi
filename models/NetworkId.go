package models

type NetworkId struct {
	Mnc string `json:"mnc,omitempty"`
	Mcc string `json:"mcc,omitempty"`
}
