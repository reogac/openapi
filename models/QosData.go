package models

type QosData struct {
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	QosId                string `json:"qosId"`
	FiveQi               *int   `json:"5qi,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
}
