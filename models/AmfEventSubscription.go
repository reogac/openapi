package models

type AmfEventSubscription struct {
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	Supi                          string        `json:"supi,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	SourceNfType                  string        `json:"sourceNfType,omitempty"`
}
