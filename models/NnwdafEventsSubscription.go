package models

type NnwdafEventsSubscription struct {
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
	EvtReq             *ReportingInformation  `json:"evtReq,omitempty"`
	NotificationURI    string                 `json:"notificationURI,omitempty"`
	NotifCorrId        string                 `json:"notifCorrId,omitempty"`
	EventSubscriptions []EventSubscription    `json:"eventSubscriptions"`
	SupportedFeatures  string                 `json:"supportedFeatures,omitempty"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
}
