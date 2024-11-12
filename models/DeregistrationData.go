package models

type DeregistrationData struct {
	AccessType       AccessType           `json:"accessType,omitempty"`
	PduSessionId     *int                 `json:"pduSessionId,omitempty"`
	NewSmfInstanceId string               `json:"newSmfInstanceId,omitempty"`
	DeregReason      DeregistrationReason `json:"deregReason"`
}
