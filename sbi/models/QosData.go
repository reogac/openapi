/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:12 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	QosId                string `json:"qosId"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
}
