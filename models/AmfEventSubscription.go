package models

type AmfEventSubscription struct {
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	NfId                          string        `json:"nfId"`
	Supi                          string        `json:"supi,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	SourceNfType                  string        `json:"sourceNfType,omitempty"`
}
