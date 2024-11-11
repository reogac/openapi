package models

type AmfEventSubscription struct {
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	NfId                          string        `json:"nfId"`
	Pei                           string        `json:"pei,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	Supi                          string        `json:"supi,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	SourceNfType                  string        `json:"sourceNfType,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
}
