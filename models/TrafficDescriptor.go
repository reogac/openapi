package models
type TrafficDescriptor struct {
	 Dnn	string	`json:"dnn,omitempty"`
	 SNssai	*Snssai	`json:"sNssai,omitempty"`
	 DddTrafficDescriptorList	[]DddTrafficDescriptor	`json:"dddTrafficDescriptorList,omitempty"`
}
