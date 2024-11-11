package models

type ExtAmfEventSubscription struct {
	Options                       *AmfEventMode `json:"options,omitempty"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	BindingInfo                   []string      `json:"bindingInfo,omitempty"`
	SubscribingNfType             string        `json:"subscribingNfType,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	EventSyncInd                  *bool         `json:"eventSyncInd,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	Supi                          string        `json:"supi,omitempty"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	SourceNfType                  string        `json:"sourceNfType,omitempty"`
}
