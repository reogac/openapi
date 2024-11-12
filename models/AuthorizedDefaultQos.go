package models

type AuthorizedDefaultQos struct {
	Arp                *Arp   `json:"arp,omitempty"`
	MaxbrDl            string `json:"maxbrDl,omitempty"`
	MaxbrUl            string `json:"maxbrUl,omitempty"`
	GbrUl              string `json:"gbrUl,omitempty"`
	GbrDl              string `json:"gbrDl,omitempty"`
	ExtMaxDataBurstVol *int   `json:"extMaxDataBurstVol,omitempty"`
	FiveQi             *int   `json:"5qi,omitempty"`
	PriorityLevel      *int   `json:"priorityLevel,omitempty"`
	AverWindow         *int   `json:"averWindow,omitempty"`
	MaxDataBurstVol    *int   `json:"maxDataBurstVol,omitempty"`
}
