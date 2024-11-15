/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NonDynamic5Qi struct {
	AverWindow            *int `json:"averWindow,omitempty"`
	MaxDataBurstVol       *int `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol    *int `json:"extMaxDataBurstVol,omitempty"`
	CnPacketDelayBudgetDl *int `json:"cnPacketDelayBudgetDl,omitempty"`
	CnPacketDelayBudgetUl *int `json:"cnPacketDelayBudgetUl,omitempty"`
	PriorityLevel         *int `json:"priorityLevel,omitempty"`
}
