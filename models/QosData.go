package models

type QosData struct {
	AverWindow           *int   `json:"averWindow,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	QosId                string `json:"qosId"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
}
