package models

type AmfEventSubscription struct {
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	Supi                          string        `json:"supi,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NfId                          string        `json:"nfId"`
	Options                       *AmfEventMode `json:"options,omitempty"`
}
