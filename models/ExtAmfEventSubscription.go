package models

type ExtAmfEventSubscription struct {
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	Pei                           string                              `json:"pei,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	NfId                          string                              `json:"nfId"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	Supi                          string                              `json:"supi,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
}
