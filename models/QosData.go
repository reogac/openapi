/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	QosId                string `json:"qosId"`
	FiveQi               *int   `json:"5qi,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
}
