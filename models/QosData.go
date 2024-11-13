package models

type QosData struct {
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	QosId                string `json:"qosId"`
	FiveQi               *int   `json:"5qi,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
}
