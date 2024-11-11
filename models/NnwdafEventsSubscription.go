package models

type NnwdafEventsSubscription struct {
	NotifCorrId        string                 `json:"notifCorrId,omitempty"`
	SupportedFeatures  string                 `json:"supportedFeatures,omitempty"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
	NotificationURI    string                 `json:"notificationURI,omitempty"`
	EvtReq             *ReportingInformation  `json:"evtReq,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
	EventSubscriptions []EventSubscription    `json:"eventSubscriptions"`
}
