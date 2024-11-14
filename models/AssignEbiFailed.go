/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AssignEbiFailed struct {
	PduSessionId  int   `json:"pduSessionId"`
	FailedArpList []Arp `json:"failedArpList,omitempty"`
}
