/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServingNfIdentity struct {
	SgsnAddr     *SgsnAddress `json:"sgsnAddr,omitempty"`
	ServNfInstId string       `json:"servNfInstId,omitempty"`
	Guami        *Guami       `json:"guami,omitempty"`
	AnGwAddr     *AnGwAddress `json:"anGwAddr,omitempty"`
}
