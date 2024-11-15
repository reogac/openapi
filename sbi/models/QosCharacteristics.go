/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:25 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosCharacteristics struct {
	ResourceType       QosResourceType `json:"resourceType"`
	PriorityLevel      int             `json:"priorityLevel"`
	PacketDelayBudget  int             `json:"packetDelayBudget"`
	PacketErrorRate    string          `json:"packetErrorRate"`
	AveragingWindow    *int            `json:"averagingWindow,omitempty"`
	MaxDataBurstVol    *int            `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol *int            `json:"extMaxDataBurstVol,omitempty"`
	FiveQi             int             `json:"5qi"`
}
