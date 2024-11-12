package models

type Dynamic5Qi struct {
	CnPacketDelayBudgetUl *int            `json:"cnPacketDelayBudgetUl,omitempty"`
	PriorityLevel         int             `json:"priorityLevel"`
	PacketDelayBudget     int             `json:"packetDelayBudget"`
	PacketErrRate         string          `json:"packetErrRate"`
	MaxDataBurstVol       *int            `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol    *int            `json:"extMaxDataBurstVol,omitempty"`
	ResourceType          QosResourceType `json:"resourceType"`
	AverWindow            *int            `json:"averWindow,omitempty"`
	ExtPacketDelBudget    *int            `json:"extPacketDelBudget,omitempty"`
	CnPacketDelayBudgetDl *int            `json:"cnPacketDelayBudgetDl,omitempty"`
}
