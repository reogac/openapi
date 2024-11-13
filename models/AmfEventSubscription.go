package models

type AmfEventSubscription struct {
	EventNotifyUri                string        `json:"eventNotifyUri"`
	Supi                          string        `json:"supi,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
}
