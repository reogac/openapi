/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	GbrUl                string `json:"gbrUl,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	QosId                string `json:"qosId"`
	FiveQi               *int   `json:"5qi,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
}
