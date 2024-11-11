package models

type Dynamic5Qi struct {
	PacketErrRate         string `json:"packetErrRate"`
	AverWindow            *int   `json:"averWindow,omitempty"`
	ExtMaxDataBurstVol    *int   `json:"extMaxDataBurstVol,omitempty"`
	CnPacketDelayBudgetDl *int   `json:"cnPacketDelayBudgetDl,omitempty"`
	CnPacketDelayBudgetUl *int   `json:"cnPacketDelayBudgetUl,omitempty"`
	ResourceType          string `json:"resourceType"`
	PriorityLevel         int    `json:"priorityLevel"`
	PacketDelayBudget     int    `json:"packetDelayBudget"`
	MaxDataBurstVol       *int   `json:"maxDataBurstVol,omitempty"`
	ExtPacketDelBudget    *int   `json:"extPacketDelBudget,omitempty"`
}
