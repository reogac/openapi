package models
type MbsSubscriptionData struct {
	 MbsAllowed	*bool	`json:"mbsAllowed,omitempty"`
	 MbsSessionIdList	[]MbsSessionId	`json:"mbsSessionIdList,omitempty"`
}
