package models

type ExtAmfEventSubscription struct {
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	NfId                          string                              `json:"nfId"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
}
