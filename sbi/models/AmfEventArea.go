/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventArea struct {
	PresenceInfo *PresenceInfo `json:"presenceInfo,omitempty"`
	LadnInfo     *LadnInfo     `json:"ladnInfo,omitempty"`
	SNssai       *Snssai       `json:"sNssai,omitempty"`
	NsiId        string        `json:"nsiId,omitempty"`
}
