package models
type RegistrationLocationInfo struct {
	 AmfInstanceId	string	`json:"amfInstanceId"`
	 Guami	*Guami	`json:"guami,omitempty"`
	 PlmnId	*PlmnId	`json:"plmnId,omitempty"`
	 VgmlcAddress	*VgmlcAddress	`json:"vgmlcAddress,omitempty"`
	 AccessTypeList	[]string	`json:"accessTypeList"`
}
