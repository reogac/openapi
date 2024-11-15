/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AdditionalMeasurement struct {
	UnexpLoc      *NetworkAreaInfo          `json:"unexpLoc,omitempty"`
	UnexpFlowTeps []IpEthFlowDescription    `json:"unexpFlowTeps,omitempty"`
	UnexpWakes    []string                  `json:"unexpWakes,omitempty"`
	DdosAttack    *AddressList              `json:"ddosAttack,omitempty"`
	WrgDest       *AddressList              `json:"wrgDest,omitempty"`
	Circums       []CircumstanceDescription `json:"circums,omitempty"`
}
