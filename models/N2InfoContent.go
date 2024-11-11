package models

type N2InfoContent struct {
	NgapMessageType *int            `json:"ngapMessageType,omitempty"`
	NgapIeType      string          `json:"ngapIeType,omitempty"`
	NgapData        RefToBinaryData `json:"ngapData"`
}
