package models

type NrppaInformation struct {
	NfId              string        `json:"nfId"`
	NrppaPdu          N2InfoContent `json:"nrppaPdu"`
	ServiceInstanceId string        `json:"serviceInstanceId,omitempty"`
}
