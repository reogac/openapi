package models

type AmfEventSubscription struct {
	Supi                          string        `json:"supi,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	SourceNfType                  string        `json:"sourceNfType,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	GroupId                       string        `json:"groupId,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
}
