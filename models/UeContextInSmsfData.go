/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextInSmsfData struct {
	SmsfInfoNon3GppAccess *SmsfInfo `json:"smsfInfoNon3GppAccess,omitempty"`
	SmsfInfo3GppAccess    *SmsfInfo `json:"smsfInfo3GppAccess,omitempty"`
}
