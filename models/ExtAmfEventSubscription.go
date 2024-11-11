package models

type ExtAmfEventSubscription struct {
	NfId                          string        `json:"nfId"`
	BindingInfo                   []string      `json:"bindingInfo,omitempty"`
	SubscribingNfType             string        `json:"subscribingNfType,omitempty"`
	EventSyncInd                  *bool         `json:"eventSyncInd,omitempty"`
	NfConsumerInfo                []string      `json:"nfConsumerInfo,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	Supi                          string        `json:"supi,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	SourceNfType                  string        `json:"sourceNfType,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
}
