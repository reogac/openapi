package models

type Dynamic5Qi struct {
	CnPacketDelayBudgetUl *int            `json:"cnPacketDelayBudgetUl,omitempty"`
	PriorityLevel         int             `json:"priorityLevel"`
	MaxDataBurstVol       *int            `json:"maxDataBurstVol,omitempty"`
	ExtPacketDelBudget    *int            `json:"extPacketDelBudget,omitempty"`
	AverWindow            *int            `json:"averWindow,omitempty"`
	ExtMaxDataBurstVol    *int            `json:"extMaxDataBurstVol,omitempty"`
	CnPacketDelayBudgetDl *int            `json:"cnPacketDelayBudgetDl,omitempty"`
	ResourceType          QosResourceType `json:"resourceType"`
	PacketDelayBudget     int             `json:"packetDelayBudget"`
	PacketErrRate         string          `json:"packetErrRate"`
}
