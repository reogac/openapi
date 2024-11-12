package models

type N1MessageContainer struct {
	ServiceInstanceId string          `json:"serviceInstanceId,omitempty"`
	N1MessageClass    string          `json:"n1MessageClass"`
	N1MessageContent  RefToBinaryData `json:"n1MessageContent"`
	NfId              string          `json:"nfId,omitempty"`
}
