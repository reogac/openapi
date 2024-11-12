package models

type PduSession struct {
	PlmnId        PlmnId  `json:"plmnId"`
	SingleNssai   *Snssai `json:"singleNssai,omitempty"`
	Dnn           string  `json:"dnn"`
	SmfInstanceId string  `json:"smfInstanceId"`
}
