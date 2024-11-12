package models

type QosData struct {
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	QosId                string `json:"qosId"`
	GbrUl                string `json:"gbrUl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
}
