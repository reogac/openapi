package models
type ExtAmfEventSubscription struct {
	 Supi	string	`json:"supi,omitempty"`
	 ExcludeGpsiList	[]string	`json:"excludeGpsiList,omitempty"`
	 IncludeGpsiList	[]string	`json:"includeGpsiList,omitempty"`
	 Pei	string	`json:"pei,omitempty"`
	 NfId	string	`json:"nfId"`
	 AnyUE	*bool	`json:"anyUE,omitempty"`
	 AoiStateList	map[string]AreaOfInterestEventState	`json:"aoiStateList,omitempty"`
	 BindingInfo	[]string	`json:"bindingInfo,omitempty"`
	 EventSyncInd	*bool	`json:"eventSyncInd,omitempty"`
	 SubsChangeNotifyUri	string	`json:"subsChangeNotifyUri,omitempty"`
	 SourceNfType	NFType	`json:"sourceNfType,omitempty"`
	 EventList	[]AmfEvent	`json:"eventList"`
	 Options	*AmfEventMode	`json:"options,omitempty"`
	 NotifyCorrelationId	string	`json:"notifyCorrelationId"`
	 SubscribingNfType	NFType	`json:"subscribingNfType,omitempty"`
	 NfConsumerInfo	[]string	`json:"nfConsumerInfo,omitempty"`
	 Gpsi	string	`json:"gpsi,omitempty"`
	 EventNotifyUri	string	`json:"eventNotifyUri"`
	 IncludeSupiList	[]string	`json:"includeSupiList,omitempty"`
	 GroupId	string	`json:"groupId,omitempty"`
	 SubsChangeNotifyCorrelationId	string	`json:"subsChangeNotifyCorrelationId,omitempty"`
	 ExcludeSupiList	[]string	`json:"excludeSupiList,omitempty"`
}
