package models

type NnwdafEventsSubscription struct {
	NotifCorrId        string                 `json:"notifCorrId,omitempty"`
	SupportedFeatures  string                 `json:"supportedFeatures,omitempty"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	EventSubscriptions []EventSubscription    `json:"eventSubscriptions"`
	EvtReq             *ReportingInformation  `json:"evtReq,omitempty"`
	NotificationURI    string                 `json:"notificationURI,omitempty"`
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
}
