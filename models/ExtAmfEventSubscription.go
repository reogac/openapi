package models

type ExtAmfEventSubscription struct {
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	SubscribingNfType             string                              `json:"subscribingNfType,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	SourceNfType                  string                              `json:"sourceNfType,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	NfId                          string                              `json:"nfId"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
}
