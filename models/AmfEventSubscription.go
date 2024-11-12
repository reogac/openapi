package models

type AmfEventSubscription struct {
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	Supi                          string        `json:"supi,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	GroupId                       string        `json:"groupId,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NfId                          string        `json:"nfId"`
}
