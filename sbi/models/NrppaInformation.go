/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NrppaInformation struct {
	NfId              string        `json:"nfId"`
	NrppaPdu          N2InfoContent `json:"nrppaPdu"`
	ServiceInstanceId string        `json:"serviceInstanceId,omitempty"`
}
