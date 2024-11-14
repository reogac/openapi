package models
type Dynamic5Qi struct {
	 CnPacketDelayBudgetUl	*int	`json:"cnPacketDelayBudgetUl,omitempty"`
	 ResourceType	QosResourceType	`json:"resourceType"`
	 PriorityLevel	int	`json:"priorityLevel"`
	 AverWindow	*int	`json:"averWindow,omitempty"`
	 MaxDataBurstVol	*int	`json:"maxDataBurstVol,omitempty"`
	 ExtMaxDataBurstVol	*int	`json:"extMaxDataBurstVol,omitempty"`
	 CnPacketDelayBudgetDl	*int	`json:"cnPacketDelayBudgetDl,omitempty"`
	 PacketDelayBudget	int	`json:"packetDelayBudget"`
	 PacketErrRate	string	`json:"packetErrRate"`
	 ExtPacketDelBudget	*int	`json:"extPacketDelBudget,omitempty"`
}
