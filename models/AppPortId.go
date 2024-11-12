package models

type AppPortId struct {
	DestinationPort *int `json:"destinationPort,omitempty"`
	OriginatorPort  *int `json:"originatorPort,omitempty"`
}
