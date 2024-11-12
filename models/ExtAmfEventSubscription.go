package models

type ExtAmfEventSubscription struct {
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	NfId                          string                              `json:"nfId"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
}
