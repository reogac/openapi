package models

type N1N2MsgTxfrErrDetail struct {
	HighestPrioArp *Arp `json:"highestPrioArp,omitempty"`
	MaxWaitingTime *int `json:"maxWaitingTime,omitempty"`
	RetryAfter     *int `json:"retryAfter,omitempty"`
}
