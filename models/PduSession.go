package models
type PduSession struct {
	 Dnn	string	`json:"dnn"`
	 SmfInstanceId	string	`json:"smfInstanceId"`
	 PlmnId	PlmnId	`json:"plmnId"`
	 SingleNssai	*Snssai	`json:"singleNssai,omitempty"`
}
