package models
type QosData struct {
	 Qnc	*bool	`json:"qnc,omitempty"`
	 AverWindow	*int	`json:"averWindow,omitempty"`
	 SharingKeyDl	string	`json:"sharingKeyDl,omitempty"`
	 SharingKeyUl	string	`json:"sharingKeyUl,omitempty"`
	 DefQosFlowIndication	*bool	`json:"defQosFlowIndication,omitempty"`
	 PacketErrorRate	string	`json:"packetErrorRate,omitempty"`
	 FiveQi	*int	`json:"5qi,omitempty"`
	 MaxbrUl	string	`json:"maxbrUl,omitempty"`
	 MaxPacketLossRateDl	*int	`json:"maxPacketLossRateDl,omitempty"`
	 ExtMaxDataBurstVol	*int	`json:"extMaxDataBurstVol,omitempty"`
	 ReflectiveQos	*bool	`json:"reflectiveQos,omitempty"`
	 QosId	string	`json:"qosId"`
	 MaxbrDl	string	`json:"maxbrDl,omitempty"`
	 GbrUl	string	`json:"gbrUl,omitempty"`
	 GbrDl	string	`json:"gbrDl,omitempty"`
	 Arp	*Arp	`json:"arp,omitempty"`
	 PriorityLevel	*int	`json:"priorityLevel,omitempty"`
	 MaxDataBurstVol	*int	`json:"maxDataBurstVol,omitempty"`
	 MaxPacketLossRateUl	*int	`json:"maxPacketLossRateUl,omitempty"`
	 PacketDelayBudget	*int	`json:"packetDelayBudget,omitempty"`
}
