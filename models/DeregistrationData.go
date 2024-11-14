package models
type DeregistrationData struct {
	 PduSessionId	*int	`json:"pduSessionId,omitempty"`
	 NewSmfInstanceId	string	`json:"newSmfInstanceId,omitempty"`
	 DeregReason	DeregistrationReason	`json:"deregReason"`
	 AccessType	AccessType	`json:"accessType,omitempty"`
}
