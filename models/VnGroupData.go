package models

type VnGroupData struct {
	Dnn             string           `json:"dnn,omitempty"`
	SingleNssai     *Snssai          `json:"singleNssai,omitempty"`
	AppDescriptors  []AppDescriptor  `json:"appDescriptors,omitempty"`
	PduSessionTypes *PduSessionTypes `json:"pduSessionTypes,omitempty"`
}
