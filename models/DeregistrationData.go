/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DeregistrationData struct {
	DeregReason      DeregistrationReason `json:"deregReason"`
	AccessType       AccessType           `json:"accessType,omitempty"`
	PduSessionId     *int                 `json:"pduSessionId,omitempty"`
	NewSmfInstanceId string               `json:"newSmfInstanceId,omitempty"`
}
