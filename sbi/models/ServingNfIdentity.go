/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServingNfIdentity struct {
	ServNfInstId string       `json:"servNfInstId,omitempty"`
	Guami        *Guami       `json:"guami,omitempty"`
	AnGwAddr     *AnGwAddress `json:"anGwAddr,omitempty"`
	SgsnAddr     *SgsnAddress `json:"sgsnAddr,omitempty"`
}
