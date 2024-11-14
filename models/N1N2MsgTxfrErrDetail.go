package models
type N1N2MsgTxfrErrDetail struct {
	 RetryAfter	*int	`json:"retryAfter,omitempty"`
	 HighestPrioArp	*Arp	`json:"highestPrioArp,omitempty"`
	 MaxWaitingTime	*int	`json:"maxWaitingTime,omitempty"`
}
