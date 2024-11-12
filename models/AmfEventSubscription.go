package models

type AmfEventSubscription struct {
	EventList                     []AmfEvent    `json:"eventList"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	Supi                          string        `json:"supi,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	SourceNfType                  string        `json:"sourceNfType,omitempty"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
}
