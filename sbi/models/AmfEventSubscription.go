/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	EventList                     []AmfEvent    `json:"eventList"`
	Supi                          string        `json:"supi,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
}
