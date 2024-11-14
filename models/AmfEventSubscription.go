/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	Supi                          string        `json:"supi,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	GroupId                       string        `json:"groupId,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
}
