package models

type SubscribedDefaultQos struct {
	PriorityLevel *int `json:"priorityLevel,omitempty"`
	FiveQi        int  `json:"5qi"`
	Arp           Arp  `json:"arp"`
}
