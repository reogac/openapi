/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NnwdafEventsSubscription struct {
	EvtReq             *ReportingInformation  `json:"evtReq,omitempty"`
	NotifCorrId        string                 `json:"notifCorrId,omitempty"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
	EventSubscriptions []EventSubscription    `json:"eventSubscriptions"`
	SupportedFeatures  string                 `json:"supportedFeatures,omitempty"`
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
	NotificationURI    string                 `json:"notificationURI,omitempty"`
}
