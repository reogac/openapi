/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:38 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HssAuthenticationVectors struct {
	AvEpsAka       []AvEpsAka       `json:"AvEpsAka,omitempty"`
	AvImsGbaEapAka []AvImsGbaEapAka `json:"AvImsGbaEapAka,omitempty"`
	AvEapAkaPrime  []AvEapAkaPrime  `json:"AvEapAkaPrime,omitempty"`
}
