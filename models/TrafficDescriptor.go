package models

type TrafficDescriptor struct {
	SNssai                   *Snssai                `json:"sNssai,omitempty"`
	DddTrafficDescriptorList []DddTrafficDescriptor `json:"dddTrafficDescriptorList,omitempty"`
	Dnn                      string                 `json:"dnn,omitempty"`
}
