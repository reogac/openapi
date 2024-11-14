package models
type AuthorizedDefaultQos struct {
	 AverWindow	*int	`json:"averWindow,omitempty"`
	 MaxDataBurstVol	*int	`json:"maxDataBurstVol,omitempty"`
	 MaxbrUl	string	`json:"maxbrUl,omitempty"`
	 MaxbrDl	string	`json:"maxbrDl,omitempty"`
	 GbrUl	string	`json:"gbrUl,omitempty"`
	 FiveQi	*int	`json:"5qi,omitempty"`
	 PriorityLevel	*int	`json:"priorityLevel,omitempty"`
	 GbrDl	string	`json:"gbrDl,omitempty"`
	 ExtMaxDataBurstVol	*int	`json:"extMaxDataBurstVol,omitempty"`
	 Arp	*Arp	`json:"arp,omitempty"`
}
