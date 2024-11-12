package models

type QosCharacteristics struct {
	MaxDataBurstVol    *int            `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol *int            `json:"extMaxDataBurstVol,omitempty"`
	FiveQi             int             `json:"5qi"`
	ResourceType       QosResourceType `json:"resourceType"`
	PriorityLevel      int             `json:"priorityLevel"`
	PacketDelayBudget  int             `json:"packetDelayBudget"`
	PacketErrorRate    string          `json:"packetErrorRate"`
	AveragingWindow    *int            `json:"averagingWindow,omitempty"`
}
