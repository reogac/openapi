package models

type UpuData struct {
	SecPacket        string   `json:"secPacket,omitempty"`
	DefaultConfNssai []Snssai `json:"defaultConfNssai,omitempty"`
	RoutingId        string   `json:"routingId,omitempty"`
	Drei             *bool    `json:"drei,omitempty"`
	Aol              *bool    `json:"aol,omitempty"`
}
