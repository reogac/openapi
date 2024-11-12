package models

type AtsssCapability struct {
	Mptcp         *bool `json:"mptcp,omitempty"`
	RttWithoutPmf *bool `json:"rttWithoutPmf,omitempty"`
	AtsssLL       *bool `json:"atsssLL,omitempty"`
}
