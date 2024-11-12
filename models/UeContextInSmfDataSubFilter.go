package models

type UeContextInSmfDataSubFilter struct {
	SnssaiList   []Snssai `json:"snssaiList,omitempty"`
	EmergencyInd *bool    `json:"emergencyInd,omitempty"`
	DnnList      []string `json:"dnnList,omitempty"`
}
