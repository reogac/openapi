/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	EventList                     []AmfEvent    `json:"eventList"`
	Supi                          string        `json:"supi,omitempty"`
}
