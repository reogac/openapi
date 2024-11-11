package models

type AssignEbiData struct {
	ArpList         []Arp  `json:"arpList,omitempty"`
	ReleasedEbiList []int  `json:"releasedEbiList,omitempty"`
	OldGuami        *Guami `json:"oldGuami,omitempty"`
	PduSessionId    int    `json:"pduSessionId"`
}
