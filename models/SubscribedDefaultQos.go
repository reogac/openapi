/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscribedDefaultQos struct {
	PriorityLevel *int `json:"priorityLevel,omitempty"`
	FiveQi        int  `json:"5qi"`
	Arp           Arp  `json:"arp"`
}
