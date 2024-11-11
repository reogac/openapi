package models

type N1N2MsgTxfrFailureNotification struct {
	Cause          string `json:"cause"`
	N1n2MsgDataUri string `json:"n1n2MsgDataUri"`
}
