package models
type AuthorizedDefaultQos struct {
	 MaxDataBurstVol	*int	`json:"maxDataBurstVol,omitempty"`
	 MaxbrDl	string	`json:"maxbrDl,omitempty"`
	 GbrDl	string	`json:"gbrDl,omitempty"`
	 PriorityLevel	*int	`json:"priorityLevel,omitempty"`
	 Arp	*Arp	`json:"arp,omitempty"`
	 AverWindow	*int	`json:"averWindow,omitempty"`
	 MaxbrUl	string	`json:"maxbrUl,omitempty"`
	 GbrUl	string	`json:"gbrUl,omitempty"`
	 ExtMaxDataBurstVol	*int	`json:"extMaxDataBurstVol,omitempty"`
	 FiveQi	*int	`json:"5qi,omitempty"`
}
