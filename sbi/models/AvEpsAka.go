/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:09 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AvEpsAka struct {
	AvType HssAvType `json:"avType"`
	Rand   string    `json:"rand"`
	Xres   string    `json:"xres"`
	Autn   string    `json:"autn"`
	Kasme  string    `json:"kasme"`
}