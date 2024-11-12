package models

type QosData struct {
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	QosId                string `json:"qosId"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
}
