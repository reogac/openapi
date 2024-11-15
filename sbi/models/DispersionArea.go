/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionArea struct {
	TaiList  []Tai  `json:"taiList,omitempty"`
	NcgiList []Ncgi `json:"ncgiList,omitempty"`
	EcgiList []Ecgi `json:"ecgiList,omitempty"`
	N3gaInd  *bool  `json:"n3gaInd,omitempty"`
}
