/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AvEapAkaPrime struct {
	IkPrime string `json:"ikPrime"`
	AvType  AvType `json:"avType"`
	Rand    string `json:"rand"`
	Xres    string `json:"xres"`
	Autn    string `json:"autn"`
	CkPrime string `json:"ckPrime"`
}
