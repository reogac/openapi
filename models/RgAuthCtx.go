/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:45 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RgAuthCtx struct {
	Supi       string     `json:"supi,omitempty"`
	AuthInd    *bool      `json:"authInd,omitempty"`
	AuthResult AuthResult `json:"authResult"`
}
