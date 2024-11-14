package models
type Dynamic5Qi struct {
	 AverWindow	*int	`json:"averWindow,omitempty"`
	 CnPacketDelayBudgetUl	*int	`json:"cnPacketDelayBudgetUl,omitempty"`
	 ResourceType	QosResourceType	`json:"resourceType"`
	 PacketErrRate	string	`json:"packetErrRate"`
	 MaxDataBurstVol	*int	`json:"maxDataBurstVol,omitempty"`
	 ExtMaxDataBurstVol	*int	`json:"extMaxDataBurstVol,omitempty"`
	 ExtPacketDelBudget	*int	`json:"extPacketDelBudget,omitempty"`
	 CnPacketDelayBudgetDl	*int	`json:"cnPacketDelayBudgetDl,omitempty"`
	 PriorityLevel	int	`json:"priorityLevel"`
	 PacketDelayBudget	int	`json:"packetDelayBudget"`
}
