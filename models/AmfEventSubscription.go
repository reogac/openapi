package models

type AmfEventSubscription struct {
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	NfId                          string        `json:"nfId"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	Supi                          string        `json:"supi,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
}
