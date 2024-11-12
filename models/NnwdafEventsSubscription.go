package models
type NnwdafEventsSubscription struct {
	 EventSubscriptions	[]EventSubscription	`json:"eventSubscriptions"`
	 FailEventReports	[]FailureEventInfo	`json:"failEventReports,omitempty"`
	 PrevSub	*PrevSubInfo	`json:"prevSub,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 EventNotifications	[]EventNotification	`json:"eventNotifications,omitempty"`
	 ConsNfInfo	*ConsumerNfInformation	`json:"consNfInfo,omitempty"`
	 EvtReq	*ReportingInformation	`json:"evtReq,omitempty"`
	 NotificationURI	string	`json:"notificationURI,omitempty"`
	 NotifCorrId	string	`json:"notifCorrId,omitempty"`
}
