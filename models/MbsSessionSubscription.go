package models
type MbsSessionSubscription struct {
	 MbsSessionSubscUri	string	`json:"mbsSessionSubscUri,omitempty"`
	 MbsSessionId	*MbsSessionId	`json:"mbsSessionId,omitempty"`
	 AreaSessionId	*int	`json:"areaSessionId,omitempty"`
	 EventList	[]MbsSessionEvent	`json:"eventList"`
	 NotifyUri	string	`json:"notifyUri"`
	 NotifyCorrelationId	string	`json:"notifyCorrelationId,omitempty"`
	 ExpiryTime	string	`json:"expiryTime,omitempty"`
	 NfcInstanceId	string	`json:"nfcInstanceId,omitempty"`
}
