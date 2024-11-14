/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:02 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConfirmationDataResponse struct {
	AuthResult AuthResult             `json:"authResult"`
	Supi       string                 `json:"supi,omitempty"`
	Kseaf      string                 `json:"kseaf,omitempty"`
	PvsInfo    []ServerAddressingInfo `json:"pvsInfo,omitempty"`
}
