package models

type PlmnIdNid struct {
	Nid string `json:"nid,omitempty"`
	Mcc string `json:"mcc"`
	Mnc string `json:"mnc"`
}
