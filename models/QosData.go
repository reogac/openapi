package models

type QosData struct {
	Qnc                  *bool  `json:"qnc,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	QosId                string `json:"qosId"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
}
