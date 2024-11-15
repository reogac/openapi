/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AreaScope struct {
	NrCellIdList    []string           `json:"nrCellIdList,omitempty"`
	TacList         []string           `json:"tacList,omitempty"`
	TacInfoPerPlmn  map[string]TacInfo `json:"tacInfoPerPlmn,omitempty"`
	EutraCellIdList []string           `json:"eutraCellIdList,omitempty"`
}
