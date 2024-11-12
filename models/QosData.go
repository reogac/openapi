package models

type QosData struct {
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	QosId                string `json:"qosId"`
	Arp                  *Arp   `json:"arp,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
}
