/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VplmnQos struct {
	Arp         *Arp   `json:"arp,omitempty"`
	SessionAmbr *Ambr  `json:"sessionAmbr,omitempty"`
	MaxFbrDl    string `json:"maxFbrDl,omitempty"`
	MaxFbrUl    string `json:"maxFbrUl,omitempty"`
	GuaFbrDl    string `json:"guaFbrDl,omitempty"`
	GuaFbrUl    string `json:"guaFbrUl,omitempty"`
	FiveQi      *int   `json:"5qi,omitempty"`
}
