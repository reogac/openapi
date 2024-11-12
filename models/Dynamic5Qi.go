package models

type Dynamic5Qi struct {
	ExtPacketDelBudget    *int            `json:"extPacketDelBudget,omitempty"`
	CnPacketDelayBudgetDl *int            `json:"cnPacketDelayBudgetDl,omitempty"`
	ExtMaxDataBurstVol    *int            `json:"extMaxDataBurstVol,omitempty"`
	PriorityLevel         int             `json:"priorityLevel"`
	PacketDelayBudget     int             `json:"packetDelayBudget"`
	PacketErrRate         string          `json:"packetErrRate"`
	AverWindow            *int            `json:"averWindow,omitempty"`
	MaxDataBurstVol       *int            `json:"maxDataBurstVol,omitempty"`
	CnPacketDelayBudgetUl *int            `json:"cnPacketDelayBudgetUl,omitempty"`
	ResourceType          QosResourceType `json:"resourceType"`
}
