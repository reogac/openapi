package models

type CandidateForReplacement struct {
	Dnns   []string `json:"dnns,omitempty"`
	Snssai Snssai   `json:"snssai"`
}
