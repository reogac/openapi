/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:10 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationLocationInfo struct {
	PlmnId         *PlmnId       `json:"plmnId,omitempty"`
	VgmlcAddress   *VgmlcAddress `json:"vgmlcAddress,omitempty"`
	AccessTypeList []string      `json:"accessTypeList"`
	AmfInstanceId  string        `json:"amfInstanceId"`
	Guami          *Guami        `json:"guami,omitempty"`
}
