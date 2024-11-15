/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NnwdafEventsSubscription struct {
	EvtReq             *ReportingInformation  `json:"evtReq,omitempty"`
	NotifCorrId        string                 `json:"notifCorrId,omitempty"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	EventSubscriptions []EventSubscription    `json:"eventSubscriptions"`
	NotificationURI    string                 `json:"notificationURI,omitempty"`
	SupportedFeatures  string                 `json:"supportedFeatures,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
}
