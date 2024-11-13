package models

type ExtAmfEventSubscription struct {
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	NfId                          string                              `json:"nfId"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
}
