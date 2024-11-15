/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AssignEbiData struct {
	PduSessionId    int             `json:"pduSessionId"`
	ArpList         []Arp           `json:"arpList,omitempty"`
	ReleasedEbiList []int           `json:"releasedEbiList,omitempty"`
	OldGuami        *Guami          `json:"oldGuami,omitempty"`
	ModifiedEbiList []EbiArpMapping `json:"modifiedEbiList,omitempty"`
}
