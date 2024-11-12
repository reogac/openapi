package models

type AuthorizedDefaultQos struct {
	Arp                *Arp   `json:"arp,omitempty"`
	MaxDataBurstVol    *int   `json:"maxDataBurstVol,omitempty"`
	MaxbrDl            string `json:"maxbrDl,omitempty"`
	GbrUl              string `json:"gbrUl,omitempty"`
	GbrDl              string `json:"gbrDl,omitempty"`
	FiveQi             *int   `json:"5qi,omitempty"`
	PriorityLevel      *int   `json:"priorityLevel,omitempty"`
	AverWindow         *int   `json:"averWindow,omitempty"`
	MaxbrUl            string `json:"maxbrUl,omitempty"`
	ExtMaxDataBurstVol *int   `json:"extMaxDataBurstVol,omitempty"`
}
