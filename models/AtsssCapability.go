package models

type AtsssCapability struct {
	AtsssLL       *bool `json:"atsssLL,omitempty"`
	Mptcp         *bool `json:"mptcp,omitempty"`
	RttWithoutPmf *bool `json:"rttWithoutPmf,omitempty"`
}
