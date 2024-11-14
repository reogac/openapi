package models
type SubscribedDefaultQos struct {
	 FiveQi	int	`json:"5qi"`
	 Arp	Arp	`json:"arp"`
	 PriorityLevel	*int	`json:"priorityLevel,omitempty"`
}
