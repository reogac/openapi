package models

type ObservedRedundantTransExp struct {
	AvgPktDelayUl    *int     `json:"avgPktDelayUl,omitempty"`
	VarPktDelayUl    *float64 `json:"varPktDelayUl,omitempty"`
	AvgPktDelayDl    *int     `json:"avgPktDelayDl,omitempty"`
	VarPktDelayDl    *float64 `json:"varPktDelayDl,omitempty"`
	AvgPktDropRateUl *int     `json:"avgPktDropRateUl,omitempty"`
	VarPktDropRateUl *float64 `json:"varPktDropRateUl,omitempty"`
	AvgPktDropRateDl *int     `json:"avgPktDropRateDl,omitempty"`
	VarPktDropRateDl *float64 `json:"varPktDropRateDl,omitempty"`
}
