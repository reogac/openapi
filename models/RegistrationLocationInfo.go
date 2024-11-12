package models

type RegistrationLocationInfo struct {
	AccessTypeList []string      `json:"accessTypeList"`
	AmfInstanceId  string        `json:"amfInstanceId"`
	Guami          *Guami        `json:"guami,omitempty"`
	PlmnId         *PlmnId       `json:"plmnId,omitempty"`
	VgmlcAddress   *VgmlcAddress `json:"vgmlcAddress,omitempty"`
}
