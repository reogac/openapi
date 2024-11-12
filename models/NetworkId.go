package models

type NetworkId struct {
	Mcc string `json:"mcc,omitempty"`
	Mnc string `json:"mnc,omitempty"`
}
