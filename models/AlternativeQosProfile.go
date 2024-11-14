/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AlternativeQosProfile struct {
	GuaFbrDl          string `json:"guaFbrDl,omitempty"`
	GuaFbrUl          string `json:"guaFbrUl,omitempty"`
	PacketDelayBudget *int   `json:"packetDelayBudget,omitempty"`
	PacketErrRate     string `json:"packetErrRate,omitempty"`
	Index             int    `json:"index"`
}
