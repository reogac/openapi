package models
type TrafficDescriptor struct {
	 DddTrafficDescriptorList	[]DddTrafficDescriptor	`json:"dddTrafficDescriptorList,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
	 SNssai	*Snssai	`json:"sNssai,omitempty"`
}
