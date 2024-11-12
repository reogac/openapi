package models

type MbsSubscriptionData struct {
	MbsSessionIdList []MbsSessionId `json:"mbsSessionIdList,omitempty"`
	MbsAllowed       *bool          `json:"mbsAllowed,omitempty"`
}
