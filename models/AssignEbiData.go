/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AssignEbiData struct {
	ReleasedEbiList []int           `json:"releasedEbiList,omitempty"`
	OldGuami        *Guami          `json:"oldGuami,omitempty"`
	ModifiedEbiList []EbiArpMapping `json:"modifiedEbiList,omitempty"`
	PduSessionId    int             `json:"pduSessionId"`
	ArpList         []Arp           `json:"arpList,omitempty"`
}
