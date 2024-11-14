/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
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
