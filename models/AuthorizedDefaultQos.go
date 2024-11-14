/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthorizedDefaultQos struct {
	GbrDl              string `json:"gbrDl,omitempty"`
	ExtMaxDataBurstVol *int   `json:"extMaxDataBurstVol,omitempty"`
	FiveQi             *int   `json:"5qi,omitempty"`
	AverWindow         *int   `json:"averWindow,omitempty"`
	MaxbrUl            string `json:"maxbrUl,omitempty"`
	MaxbrDl            string `json:"maxbrDl,omitempty"`
	GbrUl              string `json:"gbrUl,omitempty"`
	Arp                *Arp   `json:"arp,omitempty"`
	PriorityLevel      *int   `json:"priorityLevel,omitempty"`
	MaxDataBurstVol    *int   `json:"maxDataBurstVol,omitempty"`
}
