/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TargetArea struct {
	AnyTa        *bool      `json:"anyTa,omitempty"`
	TaList       []Tai      `json:"taList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
}
