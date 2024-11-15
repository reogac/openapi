/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AvImsGbaEapAka struct {
	Ck     string    `json:"ck"`
	Ik     string    `json:"ik"`
	AvType HssAvType `json:"avType"`
	Rand   string    `json:"rand"`
	Xres   string    `json:"xres"`
	Autn   string    `json:"autn"`
}
