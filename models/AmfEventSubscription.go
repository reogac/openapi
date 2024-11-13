package models

type AmfEventSubscription struct {
	EventList                     []AmfEvent    `json:"eventList"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	GroupId                       string        `json:"groupId,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	NfId                          string        `json:"nfId"`
	Supi                          string        `json:"supi,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
}
