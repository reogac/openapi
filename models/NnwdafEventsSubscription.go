package models

type NnwdafEventsSubscription struct {
	NotificationURI    string                 `json:"notificationURI,omitempty"`
	SupportedFeatures  string                 `json:"supportedFeatures,omitempty"`
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
	EventSubscriptions []EventSubscription    `json:"eventSubscriptions"`
	EvtReq             *ReportingInformation  `json:"evtReq,omitempty"`
	NotifCorrId        string                 `json:"notifCorrId,omitempty"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
}
