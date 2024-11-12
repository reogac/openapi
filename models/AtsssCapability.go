package models

type AtsssCapability struct {
	RttWithoutPmf *bool `json:"rttWithoutPmf,omitempty"`
	AtsssLL       *bool `json:"atsssLL,omitempty"`
	Mptcp         *bool `json:"mptcp,omitempty"`
}
