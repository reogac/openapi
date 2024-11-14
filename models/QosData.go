/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	QosId                string `json:"qosId"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
}
