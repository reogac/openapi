package models

type Dynamic5Qi struct {
	AverWindow            *int            `json:"averWindow,omitempty"`
	ExtMaxDataBurstVol    *int            `json:"extMaxDataBurstVol,omitempty"`
	PriorityLevel         int             `json:"priorityLevel"`
	PacketDelayBudget     int             `json:"packetDelayBudget"`
	MaxDataBurstVol       *int            `json:"maxDataBurstVol,omitempty"`
	ExtPacketDelBudget    *int            `json:"extPacketDelBudget,omitempty"`
	CnPacketDelayBudgetDl *int            `json:"cnPacketDelayBudgetDl,omitempty"`
	CnPacketDelayBudgetUl *int            `json:"cnPacketDelayBudgetUl,omitempty"`
	ResourceType          QosResourceType `json:"resourceType"`
	PacketErrRate         string          `json:"packetErrRate"`
}
