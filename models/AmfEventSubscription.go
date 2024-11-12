package models

type AmfEventSubscription struct {
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	Supi                          string        `json:"supi,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
}
