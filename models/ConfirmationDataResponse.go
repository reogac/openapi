/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:45 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConfirmationDataResponse struct {
	PvsInfo    []ServerAddressingInfo `json:"pvsInfo,omitempty"`
	AuthResult AuthResult             `json:"authResult"`
	Supi       string                 `json:"supi,omitempty"`
	Kseaf      string                 `json:"kseaf,omitempty"`
}
