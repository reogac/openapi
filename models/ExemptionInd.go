package models

type ExemptionInd struct {
	SnssaiDnnCongestion  *bool `json:"snssaiDnnCongestion,omitempty"`
	DnnCongestion        *bool `json:"dnnCongestion,omitempty"`
	SnssaiOnlyCongestion *bool `json:"snssaiOnlyCongestion,omitempty"`
}
