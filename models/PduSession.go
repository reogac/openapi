package models

type PduSession struct {
	SingleNssai   *Snssai `json:"singleNssai,omitempty"`
	Dnn           string  `json:"dnn"`
	SmfInstanceId string  `json:"smfInstanceId"`
	PlmnId        PlmnId  `json:"plmnId"`
}
