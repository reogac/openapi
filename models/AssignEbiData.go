package models
type AssignEbiData struct {
	 PduSessionId	int	`json:"pduSessionId"`
	 ArpList	[]Arp	`json:"arpList,omitempty"`
	 ReleasedEbiList	[]int	`json:"releasedEbiList,omitempty"`
	 OldGuami	*Guami	`json:"oldGuami,omitempty"`
	 ModifiedEbiList	[]EbiArpMapping	`json:"modifiedEbiList,omitempty"`
}
