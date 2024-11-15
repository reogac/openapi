/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateError struct {
	Error        ProblemDetails   `json:"error"`
	N1smCause    string           `json:"n1smCause,omitempty"`
	N1SmInfoToUe *RefToBinaryData `json:"n1SmInfoToUe,omitempty"`
	BackOffTimer *int             `json:"backOffTimer,omitempty"`
	RecoveryTime string           `json:"recoveryTime,omitempty"`
}
