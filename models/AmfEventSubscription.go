package models

type AmfEventSubscription struct {
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	NfId                          string        `json:"nfId"`
	GroupId                       string        `json:"groupId,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	Supi                          string        `json:"supi,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
}
