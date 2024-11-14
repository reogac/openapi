/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	Supi                          string                              `json:"supi,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	NfId                          string                              `json:"nfId"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	EventList                     []AmfEvent                          `json:"eventList"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
}
