package models
type N2InfoContent struct {
	 NgapMessageType	*int	`json:"ngapMessageType,omitempty"`
	 NgapIeType	NgapIeType	`json:"ngapIeType,omitempty"`
	 NgapData	RefToBinaryData	`json:"ngapData"`
}
