package models

type UpuData struct {
	Aol              *bool    `json:"aol,omitempty"`
	SecPacket        string   `json:"secPacket,omitempty"`
	DefaultConfNssai []Snssai `json:"defaultConfNssai,omitempty"`
	RoutingId        string   `json:"routingId,omitempty"`
	Drei             *bool    `json:"drei,omitempty"`
}
