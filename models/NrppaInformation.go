package models

type NrppaInformation struct {
	NrppaPdu          N2InfoContent `json:"nrppaPdu"`
	ServiceInstanceId string        `json:"serviceInstanceId,omitempty"`
	NfId              string        `json:"nfId"`
}
