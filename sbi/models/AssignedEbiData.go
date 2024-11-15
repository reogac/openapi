/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AssignedEbiData struct {
	PduSessionId    int             `json:"pduSessionId"`
	AssignedEbiList []EbiArpMapping `json:"assignedEbiList"`
	FailedArpList   []Arp           `json:"failedArpList,omitempty"`
	ReleasedEbiList []int           `json:"releasedEbiList,omitempty"`
	ModifiedEbiList []int           `json:"modifiedEbiList,omitempty"`
}
