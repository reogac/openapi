/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MsgTxfrFailureNotification struct {
	Cause          N1N2MessageTransferCause `json:"cause"`
	N1n2MsgDataUri string                   `json:"n1n2MsgDataUri"`
}
