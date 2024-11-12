package models

type AuthorizedDefaultQos struct {
	GbrDl              string `json:"gbrDl,omitempty"`
	Arp                *Arp   `json:"arp,omitempty"`
	PriorityLevel      *int   `json:"priorityLevel,omitempty"`
	MaxDataBurstVol    *int   `json:"maxDataBurstVol,omitempty"`
	MaxbrDl            string `json:"maxbrDl,omitempty"`
	GbrUl              string `json:"gbrUl,omitempty"`
	FiveQi             *int   `json:"5qi,omitempty"`
	AverWindow         *int   `json:"averWindow,omitempty"`
	MaxbrUl            string `json:"maxbrUl,omitempty"`
	ExtMaxDataBurstVol *int   `json:"extMaxDataBurstVol,omitempty"`
}
