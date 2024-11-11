package models

type SACEventStatus struct {
	ReachedNumUes     *SACInfo `json:"reachedNumUes,omitempty"`
	ReachedNumPduSess *SACInfo `json:"reachedNumPduSess,omitempty"`
}
