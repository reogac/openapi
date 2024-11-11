package models

type AssignEbiData struct {
	ModifiedEbiList []EbiArpMapping `json:"modifiedEbiList,omitempty"`
	PduSessionId    int             `json:"pduSessionId"`
	ArpList         []Arp           `json:"arpList,omitempty"`
	ReleasedEbiList []int           `json:"releasedEbiList,omitempty"`
	OldGuami        *Guami          `json:"oldGuami,omitempty"`
}
