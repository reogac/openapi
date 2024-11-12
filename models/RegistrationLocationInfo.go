package models

type RegistrationLocationInfo struct {
	Guami          *Guami        `json:"guami,omitempty"`
	PlmnId         *PlmnId       `json:"plmnId,omitempty"`
	VgmlcAddress   *VgmlcAddress `json:"vgmlcAddress,omitempty"`
	AccessTypeList []string      `json:"accessTypeList"`
	AmfInstanceId  string        `json:"amfInstanceId"`
}
