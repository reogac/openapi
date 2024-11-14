package models
type DeregistrationData struct {
	 NewSmfInstanceId	string	`json:"newSmfInstanceId,omitempty"`
	 DeregReason	DeregistrationReason	`json:"deregReason"`
	 AccessType	AccessType	`json:"accessType,omitempty"`
	 PduSessionId	*int	`json:"pduSessionId,omitempty"`
}
