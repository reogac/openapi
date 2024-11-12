package models

type Dynamic5Qi struct {
	ResourceType          QosResourceType `json:"resourceType"`
	PriorityLevel         int             `json:"priorityLevel"`
	PacketErrRate         string          `json:"packetErrRate"`
	CnPacketDelayBudgetDl *int            `json:"cnPacketDelayBudgetDl,omitempty"`
	CnPacketDelayBudgetUl *int            `json:"cnPacketDelayBudgetUl,omitempty"`
	PacketDelayBudget     int             `json:"packetDelayBudget"`
	AverWindow            *int            `json:"averWindow,omitempty"`
	MaxDataBurstVol       *int            `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol    *int            `json:"extMaxDataBurstVol,omitempty"`
	ExtPacketDelBudget    *int            `json:"extPacketDelBudget,omitempty"`
}
