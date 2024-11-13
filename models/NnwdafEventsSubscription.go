package models

type NnwdafEventsSubscription struct {
	EventSubscriptions []EventSubscription    `json:"eventSubscriptions"`
	EvtReq             *ReportingInformation  `json:"evtReq,omitempty"`
	NotifCorrId        string                 `json:"notifCorrId,omitempty"`
	SupportedFeatures  string                 `json:"supportedFeatures,omitempty"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
	NotificationURI    string                 `json:"notificationURI,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
}
