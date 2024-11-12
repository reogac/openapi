package models

type Dynamic5Qi struct {
	MaxDataBurstVol       *int   `json:"maxDataBurstVol,omitempty"`
	CnPacketDelayBudgetDl *int   `json:"cnPacketDelayBudgetDl,omitempty"`
	PacketErrRate         string `json:"packetErrRate"`
	PriorityLevel         int    `json:"priorityLevel"`
	PacketDelayBudget     int    `json:"packetDelayBudget"`
	AverWindow            *int   `json:"averWindow,omitempty"`
	ExtMaxDataBurstVol    *int   `json:"extMaxDataBurstVol,omitempty"`
	ExtPacketDelBudget    *int   `json:"extPacketDelBudget,omitempty"`
	CnPacketDelayBudgetUl *int   `json:"cnPacketDelayBudgetUl,omitempty"`
	ResourceType          string `json:"resourceType"`
}
