/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AdditionalMeasurement struct {
	WrgDest       *AddressList              `json:"wrgDest,omitempty"`
	Circums       []CircumstanceDescription `json:"circums,omitempty"`
	UnexpLoc      *NetworkAreaInfo          `json:"unexpLoc,omitempty"`
	UnexpFlowTeps []IpEthFlowDescription    `json:"unexpFlowTeps,omitempty"`
	UnexpWakes    []string                  `json:"unexpWakes,omitempty"`
	DdosAttack    *AddressList              `json:"ddosAttack,omitempty"`
}
