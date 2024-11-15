/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Dynamic5Qi struct {
	CnPacketDelayBudgetUl *int            `json:"cnPacketDelayBudgetUl,omitempty"`
	ResourceType          QosResourceType `json:"resourceType"`
	PriorityLevel         int             `json:"priorityLevel"`
	PacketErrRate         string          `json:"packetErrRate"`
	MaxDataBurstVol       *int            `json:"maxDataBurstVol,omitempty"`
	CnPacketDelayBudgetDl *int            `json:"cnPacketDelayBudgetDl,omitempty"`
	PacketDelayBudget     int             `json:"packetDelayBudget"`
	AverWindow            *int            `json:"averWindow,omitempty"`
	ExtMaxDataBurstVol    *int            `json:"extMaxDataBurstVol,omitempty"`
	ExtPacketDelBudget    *int            `json:"extPacketDelBudget,omitempty"`
}
