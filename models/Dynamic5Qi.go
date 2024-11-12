package models

type Dynamic5Qi struct {
	CnPacketDelayBudgetUl *int            `json:"cnPacketDelayBudgetUl,omitempty"`
	ResourceType          QosResourceType `json:"resourceType"`
	PriorityLevel         int             `json:"priorityLevel"`
	PacketDelayBudget     int             `json:"packetDelayBudget"`
	PacketErrRate         string          `json:"packetErrRate"`
	MaxDataBurstVol       *int            `json:"maxDataBurstVol,omitempty"`
	ExtPacketDelBudget    *int            `json:"extPacketDelBudget,omitempty"`
	CnPacketDelayBudgetDl *int            `json:"cnPacketDelayBudgetDl,omitempty"`
	AverWindow            *int            `json:"averWindow,omitempty"`
	ExtMaxDataBurstVol    *int            `json:"extMaxDataBurstVol,omitempty"`
}
