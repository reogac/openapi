/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	NfId                          string                              `json:"nfId"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
}
