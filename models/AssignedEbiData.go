package models
type AssignedEbiData struct {
	 FailedArpList	[]Arp	`json:"failedArpList,omitempty"`
	 ReleasedEbiList	[]int	`json:"releasedEbiList,omitempty"`
	 ModifiedEbiList	[]int	`json:"modifiedEbiList,omitempty"`
	 PduSessionId	int	`json:"pduSessionId"`
	 AssignedEbiList	[]EbiArpMapping	`json:"assignedEbiList"`
}
