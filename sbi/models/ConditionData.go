/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:25 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConditionData struct {
	CondId           string     `json:"condId"`
	ActivationTime   string     `json:"activationTime,omitempty"`
	DeactivationTime string     `json:"deactivationTime,omitempty"`
	AccessType       AccessType `json:"accessType,omitempty"`
	RatType          RatType    `json:"ratType,omitempty"`
}
