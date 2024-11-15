/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:12 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthorizedDefaultQos struct {
	PriorityLevel      *int   `json:"priorityLevel,omitempty"`
	MaxbrDl            string `json:"maxbrDl,omitempty"`
	GbrUl              string `json:"gbrUl,omitempty"`
	GbrDl              string `json:"gbrDl,omitempty"`
	MaxbrUl            string `json:"maxbrUl,omitempty"`
	ExtMaxDataBurstVol *int   `json:"extMaxDataBurstVol,omitempty"`
	FiveQi             *int   `json:"5qi,omitempty"`
	Arp                *Arp   `json:"arp,omitempty"`
	AverWindow         *int   `json:"averWindow,omitempty"`
	MaxDataBurstVol    *int   `json:"maxDataBurstVol,omitempty"`
}
