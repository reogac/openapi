/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NnwdafEventsSubscription struct {
	EvtReq             *ReportingInformation  `json:"evtReq,omitempty"`
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	EventSubscriptions []EventSubscription    `json:"eventSubscriptions"`
	NotificationURI    string                 `json:"notificationURI,omitempty"`
	NotifCorrId        string                 `json:"notifCorrId,omitempty"`
	SupportedFeatures  string                 `json:"supportedFeatures,omitempty"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
}
