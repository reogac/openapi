package models

type SteeringContainer struct {
	SecuredPacket string         `json:"SecuredPacket,omitempty"`
	SteeringInfo  []SteeringInfo `json:"SteeringInfo,omitempty"`
}
