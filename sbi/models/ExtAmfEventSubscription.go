/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:28 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	NfId                          string                              `json:"nfId"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	Supi                          string                              `json:"supi,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
}
