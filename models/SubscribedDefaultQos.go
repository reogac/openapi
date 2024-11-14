package models
type SubscribedDefaultQos struct {
	 Arp	Arp	`json:"arp"`
	 PriorityLevel	*int	`json:"priorityLevel,omitempty"`
	 FiveQi	int	`json:"5qi"`
}
