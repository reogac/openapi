package models
type NnwdafEventsSubscription struct {
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 EventNotifications	[]EventNotification	`json:"eventNotifications,omitempty"`
	 FailEventReports	[]FailureEventInfo	`json:"failEventReports,omitempty"`
	 PrevSub	*PrevSubInfo	`json:"prevSub,omitempty"`
	 ConsNfInfo	*ConsumerNfInformation	`json:"consNfInfo,omitempty"`
	 EventSubscriptions	[]EventSubscription	`json:"eventSubscriptions"`
	 EvtReq	*ReportingInformation	`json:"evtReq,omitempty"`
	 NotificationURI	string	`json:"notificationURI,omitempty"`
	 NotifCorrId	string	`json:"notifCorrId,omitempty"`
}
