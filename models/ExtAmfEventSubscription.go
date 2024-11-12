package models

type ExtAmfEventSubscription struct {
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	Pei                           string                              `json:"pei,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	NfId                          string                              `json:"nfId"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
}
