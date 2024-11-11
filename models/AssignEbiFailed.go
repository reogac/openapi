package models
type AssignEbiFailed struct {
	 PduSessionId	int	`json:"pduSessionId"`
	 FailedArpList	[]Arp	`json:"failedArpList,omitempty"`
}
