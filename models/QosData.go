package models

type QosData struct {
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	QosId                string `json:"qosId"`
	GbrDl                string `json:"gbrDl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
}
