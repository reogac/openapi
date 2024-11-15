/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PlmnEcInfo struct {
	EcRestrictionDataWb *EcRestrictionDataWb `json:"ecRestrictionDataWb,omitempty"`
	EcRestrictionDataNb *bool                `json:"ecRestrictionDataNb,omitempty"`
	PlmnId              PlmnId               `json:"plmnId"`
}
