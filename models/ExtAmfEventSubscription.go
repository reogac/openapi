package models

type ExtAmfEventSubscription struct {
	EventList                     []AmfEvent                          `json:"eventList"`
	NfId                          string                              `json:"nfId"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	GroupId                       string                              `json:"groupId,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
}
