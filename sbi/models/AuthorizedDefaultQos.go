/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthorizedDefaultQos struct {
	PriorityLevel      *int   `json:"priorityLevel,omitempty"`
	AverWindow         *int   `json:"averWindow,omitempty"`
	GbrDl              string `json:"gbrDl,omitempty"`
	ExtMaxDataBurstVol *int   `json:"extMaxDataBurstVol,omitempty"`
	FiveQi             *int   `json:"5qi,omitempty"`
	Arp                *Arp   `json:"arp,omitempty"`
	MaxDataBurstVol    *int   `json:"maxDataBurstVol,omitempty"`
	MaxbrUl            string `json:"maxbrUl,omitempty"`
	MaxbrDl            string `json:"maxbrDl,omitempty"`
	GbrUl              string `json:"gbrUl,omitempty"`
}
