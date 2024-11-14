/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:38 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AvEpsAka struct {
	Autn   string    `json:"autn"`
	Kasme  string    `json:"kasme"`
	AvType HssAvType `json:"avType"`
	Rand   string    `json:"rand"`
	Xres   string    `json:"xres"`
}
