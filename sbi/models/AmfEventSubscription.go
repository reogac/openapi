/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	EventList                     []AmfEvent    `json:"eventList"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	Supi                          string        `json:"supi,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
}
