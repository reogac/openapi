package models
type N1MessageContainer struct {
	 N1MessageContent	RefToBinaryData	`json:"n1MessageContent"`
	 NfId	string	`json:"nfId,omitempty"`
	 ServiceInstanceId	string	`json:"serviceInstanceId,omitempty"`
	 N1MessageClass	N1MessageClass	`json:"n1MessageClass"`
}
