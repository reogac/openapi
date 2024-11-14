package models
type QosData struct {
	 MaxbrDl	string	`json:"maxbrDl,omitempty"`
	 Arp	*Arp	`json:"arp,omitempty"`
	 SharingKeyUl	string	`json:"sharingKeyUl,omitempty"`
	 PacketDelayBudget	*int	`json:"packetDelayBudget,omitempty"`
	 PacketErrorRate	string	`json:"packetErrorRate,omitempty"`
	 GbrDl	string	`json:"gbrDl,omitempty"`
	 AverWindow	*int	`json:"averWindow,omitempty"`
	 ReflectiveQos	*bool	`json:"reflectiveQos,omitempty"`
	 MaxPacketLossRateDl	*int	`json:"maxPacketLossRateDl,omitempty"`
	 DefQosFlowIndication	*bool	`json:"defQosFlowIndication,omitempty"`
	 FiveQi	*int	`json:"5qi,omitempty"`
	 MaxbrUl	string	`json:"maxbrUl,omitempty"`
	 GbrUl	string	`json:"gbrUl,omitempty"`
	 PriorityLevel	*int	`json:"priorityLevel,omitempty"`
	 MaxDataBurstVol	*int	`json:"maxDataBurstVol,omitempty"`
	 SharingKeyDl	string	`json:"sharingKeyDl,omitempty"`
	 QosId	string	`json:"qosId"`
	 Qnc	*bool	`json:"qnc,omitempty"`
	 MaxPacketLossRateUl	*int	`json:"maxPacketLossRateUl,omitempty"`
	 ExtMaxDataBurstVol	*int	`json:"extMaxDataBurstVol,omitempty"`
}
