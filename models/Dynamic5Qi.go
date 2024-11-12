package models

type Dynamic5Qi struct {
	ResourceType          string `json:"resourceType"`
	MaxDataBurstVol       *int   `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol    *int   `json:"extMaxDataBurstVol,omitempty"`
	ExtPacketDelBudget    *int   `json:"extPacketDelBudget,omitempty"`
	PriorityLevel         int    `json:"priorityLevel"`
	PacketDelayBudget     int    `json:"packetDelayBudget"`
	PacketErrRate         string `json:"packetErrRate"`
	AverWindow            *int   `json:"averWindow,omitempty"`
	CnPacketDelayBudgetDl *int   `json:"cnPacketDelayBudgetDl,omitempty"`
	CnPacketDelayBudgetUl *int   `json:"cnPacketDelayBudgetUl,omitempty"`
}
