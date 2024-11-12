package models

type ThresholdValue struct {
	RttThres *int `json:"rttThres,omitempty"`
	PlrThres *int `json:"plrThres,omitempty"`
}
