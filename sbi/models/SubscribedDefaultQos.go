/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:12 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscribedDefaultQos struct {
	FiveQi        int  `json:"5qi"`
	Arp           Arp  `json:"arp"`
	PriorityLevel *int `json:"priorityLevel,omitempty"`
}
