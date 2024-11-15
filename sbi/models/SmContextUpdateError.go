/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextUpdateError struct {
	RecoveryTime string            `json:"recoveryTime,omitempty"`
	Error        ExtProblemDetails `json:"error"`
	N1SmMsg      *RefToBinaryData  `json:"n1SmMsg,omitempty"`
	N2SmInfo     *RefToBinaryData  `json:"n2SmInfo,omitempty"`
	N2SmInfoType N2SmInfoType      `json:"n2SmInfoType,omitempty"`
	UpCnxState   UpCnxState        `json:"upCnxState,omitempty"`
}
