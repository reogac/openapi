/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:08 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextInSmsfData struct {
	SmsfInfo3GppAccess    *SmsfInfo `json:"smsfInfo3GppAccess,omitempty"`
	SmsfInfoNon3GppAccess *SmsfInfo `json:"smsfInfoNon3GppAccess,omitempty"`
}