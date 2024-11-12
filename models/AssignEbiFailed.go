package models

type AssignEbiFailed struct {
	FailedArpList []Arp `json:"failedArpList,omitempty"`
	PduSessionId  int   `json:"pduSessionId"`
}
