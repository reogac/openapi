package models

type ExtAmfEventSubscription struct {
	EventList                     []AmfEvent    `json:"eventList"`
	SubscribingNfType             string        `json:"subscribingNfType,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	SourceNfType                  string        `json:"sourceNfType,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	EventSyncInd                  *bool         `json:"eventSyncInd,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	BindingInfo                   []string      `json:"bindingInfo,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	NfId                          string        `json:"nfId"`
	Supi                          string        `json:"supi,omitempty"`
}
