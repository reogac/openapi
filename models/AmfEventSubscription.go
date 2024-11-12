package models

type AmfEventSubscription struct {
	Pei                           string        `json:"pei,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NfId                          string        `json:"nfId"`
	Supi                          string        `json:"supi,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
}
