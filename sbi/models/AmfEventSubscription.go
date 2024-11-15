/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	Supi                          string        `json:"supi,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
}
