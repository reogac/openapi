package models

type ExtAmfEventSubscription struct {
	Pei                           string                              `json:"pei,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	NfId                          string                              `json:"nfId"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	GroupId                       string                              `json:"groupId,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
}
