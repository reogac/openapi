package models

type ExtAmfEventSubscription struct {
	Pei                           string                              `json:"pei,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	NfId                          string                              `json:"nfId"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
}
