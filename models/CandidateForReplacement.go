package models

type CandidateForReplacement struct {
	Snssai Snssai   `json:"snssai"`
	Dnns   []string `json:"dnns,omitempty"`
}
