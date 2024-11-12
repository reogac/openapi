package models

type NnwdafEventsSubscription struct {
	EventSubscriptions []EventSubscription    `json:"eventSubscriptions"`
	EvtReq             *ReportingInformation  `json:"evtReq,omitempty"`
	NotifCorrId        string                 `json:"notifCorrId,omitempty"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	NotificationURI    string                 `json:"notificationURI,omitempty"`
	SupportedFeatures  string                 `json:"supportedFeatures,omitempty"`
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
}
