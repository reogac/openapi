package models
type VnGroupData struct {
	 AppDescriptors	[]AppDescriptor	`json:"appDescriptors,omitempty"`
	 PduSessionTypes	*PduSessionTypes	`json:"pduSessionTypes,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
	 SingleNssai	*Snssai	`json:"singleNssai,omitempty"`
}
