/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	NfId                          string                              `json:"nfId"`
	GroupId                       string                              `json:"groupId,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
}
