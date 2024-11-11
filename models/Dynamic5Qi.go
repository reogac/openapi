package models
type Dynamic5Qi struct {
	 ResourceType	string	`json:"resourceType"`
	 PacketErrRate	string	`json:"packetErrRate"`
	 ExtPacketDelBudget	*int	`json:"extPacketDelBudget,omitempty"`
	 CnPacketDelayBudgetDl	*int	`json:"cnPacketDelayBudgetDl,omitempty"`
	 PriorityLevel	int	`json:"priorityLevel"`
	 PacketDelayBudget	int	`json:"packetDelayBudget"`
	 AverWindow	*int	`json:"averWindow,omitempty"`
	 MaxDataBurstVol	*int	`json:"maxDataBurstVol,omitempty"`
	 ExtMaxDataBurstVol	*int	`json:"extMaxDataBurstVol,omitempty"`
	 CnPacketDelayBudgetUl	*int	`json:"cnPacketDelayBudgetUl,omitempty"`
}
