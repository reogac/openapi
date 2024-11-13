package models

type AuthorizedDefaultQos struct {
	FiveQi             *int   `json:"5qi,omitempty"`
	Arp                *Arp   `json:"arp,omitempty"`
	MaxbrDl            string `json:"maxbrDl,omitempty"`
	GbrDl              string `json:"gbrDl,omitempty"`
	ExtMaxDataBurstVol *int   `json:"extMaxDataBurstVol,omitempty"`
	PriorityLevel      *int   `json:"priorityLevel,omitempty"`
	AverWindow         *int   `json:"averWindow,omitempty"`
	MaxDataBurstVol    *int   `json:"maxDataBurstVol,omitempty"`
	MaxbrUl            string `json:"maxbrUl,omitempty"`
	GbrUl              string `json:"gbrUl,omitempty"`
}
