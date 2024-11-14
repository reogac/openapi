package models
type VnGroupData struct {
	 SingleNssai	*Snssai	`json:"singleNssai,omitempty"`
	 AppDescriptors	[]AppDescriptor	`json:"appDescriptors,omitempty"`
	 PduSessionTypes	*PduSessionTypes	`json:"pduSessionTypes,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
}
