package models

type FiveGAuthData struct {
	EapPayload string   `json:"EapPayload,omitempty"`
	Av5gAka    *Av5gAka `json:"Av5gAka,omitempty"`
}
