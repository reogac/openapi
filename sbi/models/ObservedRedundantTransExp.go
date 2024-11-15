/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

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
