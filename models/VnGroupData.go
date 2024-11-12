package models

type VnGroupData struct {
	PduSessionTypes *PduSessionTypes `json:"pduSessionTypes,omitempty"`
	Dnn             string           `json:"dnn,omitempty"`
	SingleNssai     *Snssai          `json:"singleNssai,omitempty"`
	AppDescriptors  []AppDescriptor  `json:"appDescriptors,omitempty"`
}
