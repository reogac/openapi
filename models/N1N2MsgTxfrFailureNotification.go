package models
type N1N2MsgTxfrFailureNotification struct {
	 N1n2MsgDataUri	string	`json:"n1n2MsgDataUri"`
	 Cause	N1N2MessageTransferCause	`json:"cause"`
}
