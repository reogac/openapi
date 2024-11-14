package models
type ExtAmfEventSubscription struct {
	 BindingInfo	[]string	`json:"bindingInfo,omitempty"`
	 Pei	string	`json:"pei,omitempty"`
	 EventNotifyUri	string	`json:"eventNotifyUri"`
	 SourceNfType	NFType	`json:"sourceNfType,omitempty"`
	 NfConsumerInfo	[]string	`json:"nfConsumerInfo,omitempty"`
	 Gpsi	string	`json:"gpsi,omitempty"`
	 ExcludeSupiList	[]string	`json:"excludeSupiList,omitempty"`
	 ExcludeGpsiList	[]string	`json:"excludeGpsiList,omitempty"`
	 IncludeGpsiList	[]string	`json:"includeGpsiList,omitempty"`
	 Options	*AmfEventMode	`json:"options,omitempty"`
	 NfId	string	`json:"nfId"`
	 Supi	string	`json:"supi,omitempty"`
	 IncludeSupiList	[]string	`json:"includeSupiList,omitempty"`
	 AoiStateList	map[string]AreaOfInterestEventState	`json:"aoiStateList,omitempty"`
	 SubscribingNfType	NFType	`json:"subscribingNfType,omitempty"`
	 EventSyncInd	*bool	`json:"eventSyncInd,omitempty"`
	 GroupId	string	`json:"groupId,omitempty"`
	 EventList	[]AmfEvent	`json:"eventList"`
	 NotifyCorrelationId	string	`json:"notifyCorrelationId"`
	 SubsChangeNotifyUri	string	`json:"subsChangeNotifyUri,omitempty"`
	 AnyUE	*bool	`json:"anyUE,omitempty"`
	 SubsChangeNotifyCorrelationId	string	`json:"subsChangeNotifyCorrelationId,omitempty"`
}
