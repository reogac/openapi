package models
type UpuData struct {
	 Drei	*bool	`json:"drei,omitempty"`
	 Aol	*bool	`json:"aol,omitempty"`
	 SecPacket	string	`json:"secPacket,omitempty"`
	 DefaultConfNssai	[]Snssai	`json:"defaultConfNssai,omitempty"`
	 RoutingId	string	`json:"routingId,omitempty"`
}
