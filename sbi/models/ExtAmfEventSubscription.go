/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	NfId                          string                              `json:"nfId"`
	Supi                          string                              `json:"supi,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	GroupId                       string                              `json:"groupId,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
}
