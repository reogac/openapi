/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:08 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AdditionalSnssaiData struct {
	RequiredAuthnAuthz   *bool     `json:"requiredAuthnAuthz,omitempty"`
	SubscribedUeSliceMbr *SliceMbr `json:"subscribedUeSliceMbr,omitempty"`
	SubscribedNsSrgList  []string  `json:"subscribedNsSrgList,omitempty"`
}