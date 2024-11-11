package models

type ExtAmfEventSubscription struct {
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	NfId                          string        `json:"nfId"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	Supi                          string        `json:"supi,omitempty"`
	SourceNfType                  string        `json:"sourceNfType,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	GroupId                       string        `json:"groupId,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	NfConsumerInfo                []string      `json:"nfConsumerInfo,omitempty"`
	EventSyncInd                  *bool         `json:"eventSyncInd,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	SubscribingNfType             string        `json:"subscribingNfType,omitempty"`
	BindingInfo                   []string      `json:"bindingInfo,omitempty"`
}
