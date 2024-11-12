package models

type N2InfoContent struct {
	NgapData        RefToBinaryData `json:"ngapData"`
	NgapMessageType *int            `json:"ngapMessageType,omitempty"`
	NgapIeType      NgapIeType      `json:"ngapIeType,omitempty"`
}
