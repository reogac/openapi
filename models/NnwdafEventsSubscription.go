package models
type NnwdafEventsSubscription struct {
	 EventSubscriptions	[]EventSubscription	`json:"eventSubscriptions"`
	 NotifCorrId	string	`json:"notifCorrId,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 FailEventReports	[]FailureEventInfo	`json:"failEventReports,omitempty"`
	 PrevSub	*PrevSubInfo	`json:"prevSub,omitempty"`
	 ConsNfInfo	*ConsumerNfInformation	`json:"consNfInfo,omitempty"`
	 EvtReq	*ReportingInformation	`json:"evtReq,omitempty"`
	 NotificationURI	string	`json:"notificationURI,omitempty"`
	 EventNotifications	[]EventNotification	`json:"eventNotifications,omitempty"`
}
