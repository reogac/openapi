/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSession struct {
	Dnn           string  `json:"dnn"`
	SmfInstanceId string  `json:"smfInstanceId"`
	PlmnId        PlmnId  `json:"plmnId"`
	SingleNssai   *Snssai `json:"singleNssai,omitempty"`
}
