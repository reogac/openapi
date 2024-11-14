/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	GroupId                       string                              `json:"groupId,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	NfId                          string                              `json:"nfId"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
}
