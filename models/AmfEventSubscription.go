package models

type AmfEventSubscription struct {
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	Supi                          string        `json:"supi,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	GroupId                       string        `json:"groupId,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
}
