package models

type Dynamic5Qi struct {
	ExtPacketDelBudget    *int   `json:"extPacketDelBudget,omitempty"`
	CnPacketDelayBudgetDl *int   `json:"cnPacketDelayBudgetDl,omitempty"`
	CnPacketDelayBudgetUl *int   `json:"cnPacketDelayBudgetUl,omitempty"`
	ResourceType          string `json:"resourceType"`
	PriorityLevel         int    `json:"priorityLevel"`
	PacketDelayBudget     int    `json:"packetDelayBudget"`
	PacketErrRate         string `json:"packetErrRate"`
	AverWindow            *int   `json:"averWindow,omitempty"`
	MaxDataBurstVol       *int   `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol    *int   `json:"extMaxDataBurstVol,omitempty"`
}
