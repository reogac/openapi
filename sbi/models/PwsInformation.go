/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PwsInformation struct {
	BcEmptyAreaList   []GlobalRanNodeId `json:"bcEmptyAreaList,omitempty"`
	SendRanResponse   *bool             `json:"sendRanResponse,omitempty"`
	OmcId             string            `json:"omcId,omitempty"`
	NfId              string            `json:"nfId,omitempty"`
	MessageIdentifier int               `json:"messageIdentifier"`
	SerialNumber      int               `json:"serialNumber"`
	PwsContainer      N2InfoContent     `json:"pwsContainer"`
}
