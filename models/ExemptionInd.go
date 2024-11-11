package models

type ExemptionInd struct {
	SnssaiOnlyCongestion *bool `json:"snssaiOnlyCongestion,omitempty"`
	SnssaiDnnCongestion  *bool `json:"snssaiDnnCongestion,omitempty"`
	DnnCongestion        *bool `json:"dnnCongestion,omitempty"`
}
