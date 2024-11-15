/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	Arp                  *Arp   `json:"arp,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	QosId                string `json:"qosId"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	Qnc                  *bool  `json:"qnc,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
}
