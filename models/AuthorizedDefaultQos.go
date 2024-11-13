package models

type AuthorizedDefaultQos struct {
	Arp                *Arp   `json:"arp,omitempty"`
	PriorityLevel      *int   `json:"priorityLevel,omitempty"`
	MaxDataBurstVol    *int   `json:"maxDataBurstVol,omitempty"`
	FiveQi             *int   `json:"5qi,omitempty"`
	AverWindow         *int   `json:"averWindow,omitempty"`
	MaxbrUl            string `json:"maxbrUl,omitempty"`
	MaxbrDl            string `json:"maxbrDl,omitempty"`
	GbrUl              string `json:"gbrUl,omitempty"`
	GbrDl              string `json:"gbrDl,omitempty"`
	ExtMaxDataBurstVol *int   `json:"extMaxDataBurstVol,omitempty"`
}
