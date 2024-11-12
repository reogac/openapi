package models

type RegistrationLocationInfo struct {
	VgmlcAddress   *VgmlcAddress `json:"vgmlcAddress,omitempty"`
	AccessTypeList []string      `json:"accessTypeList"`
	AmfInstanceId  string        `json:"amfInstanceId"`
	Guami          *Guami        `json:"guami,omitempty"`
	PlmnId         *PlmnId       `json:"plmnId,omitempty"`
}
