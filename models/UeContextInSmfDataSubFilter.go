/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextInSmfDataSubFilter struct {
	EmergencyInd *bool    `json:"emergencyInd,omitempty"`
	DnnList      []string `json:"dnnList,omitempty"`
	SnssaiList   []Snssai `json:"snssaiList,omitempty"`
}
