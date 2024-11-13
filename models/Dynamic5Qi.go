package models

type Dynamic5Qi struct {
	ResourceType          QosResourceType `json:"resourceType"`
	PriorityLevel         int             `json:"priorityLevel"`
	PacketDelayBudget     int             `json:"packetDelayBudget"`
	MaxDataBurstVol       *int            `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol    *int            `json:"extMaxDataBurstVol,omitempty"`
	CnPacketDelayBudgetUl *int            `json:"cnPacketDelayBudgetUl,omitempty"`
	PacketErrRate         string          `json:"packetErrRate"`
	AverWindow            *int            `json:"averWindow,omitempty"`
	ExtPacketDelBudget    *int            `json:"extPacketDelBudget,omitempty"`
	CnPacketDelayBudgetDl *int            `json:"cnPacketDelayBudgetDl,omitempty"`
}
