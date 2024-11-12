package models

type NnwdafEventsSubscription struct {
	EventSubscriptions []EventSubscription    `json:"eventSubscriptions"`
	NotificationURI    string                 `json:"notificationURI,omitempty"`
	SupportedFeatures  string                 `json:"supportedFeatures,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	EvtReq             *ReportingInformation  `json:"evtReq,omitempty"`
	NotifCorrId        string                 `json:"notifCorrId,omitempty"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
}
