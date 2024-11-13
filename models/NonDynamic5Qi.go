package models

type NonDynamic5Qi struct {
	CnPacketDelayBudgetDl *int `json:"cnPacketDelayBudgetDl,omitempty"`
	CnPacketDelayBudgetUl *int `json:"cnPacketDelayBudgetUl,omitempty"`
	PriorityLevel         *int `json:"priorityLevel,omitempty"`
	AverWindow            *int `json:"averWindow,omitempty"`
	MaxDataBurstVol       *int `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol    *int `json:"extMaxDataBurstVol,omitempty"`
}
