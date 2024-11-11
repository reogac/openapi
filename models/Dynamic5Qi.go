package models

type Dynamic5Qi struct {
	PacketDelayBudget     int    `json:"packetDelayBudget"`
	AverWindow            *int   `json:"averWindow,omitempty"`
	MaxDataBurstVol       *int   `json:"maxDataBurstVol,omitempty"`
	CnPacketDelayBudgetUl *int   `json:"cnPacketDelayBudgetUl,omitempty"`
	ResourceType          string `json:"resourceType"`
	PriorityLevel         int    `json:"priorityLevel"`
	PacketErrRate         string `json:"packetErrRate"`
	ExtMaxDataBurstVol    *int   `json:"extMaxDataBurstVol,omitempty"`
	ExtPacketDelBudget    *int   `json:"extPacketDelBudget,omitempty"`
	CnPacketDelayBudgetDl *int   `json:"cnPacketDelayBudgetDl,omitempty"`
}
