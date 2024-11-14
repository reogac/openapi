package models
type ExtAmfEventSubscription struct {
	 IncludeSupiList	[]string	`json:"includeSupiList,omitempty"`
	 SubsChangeNotifyCorrelationId	string	`json:"subsChangeNotifyCorrelationId,omitempty"`
	 Supi	string	`json:"supi,omitempty"`
	 Options	*AmfEventMode	`json:"options,omitempty"`
	 NfId	string	`json:"nfId"`
	 IncludeGpsiList	[]string	`json:"includeGpsiList,omitempty"`
	 AnyUE	*bool	`json:"anyUE,omitempty"`
	 SubsChangeNotifyUri	string	`json:"subsChangeNotifyUri,omitempty"`
	 SourceNfType	NFType	`json:"sourceNfType,omitempty"`
	 EventNotifyUri	string	`json:"eventNotifyUri"`
	 NfConsumerInfo	[]string	`json:"nfConsumerInfo,omitempty"`
	 ExcludeSupiList	[]string	`json:"excludeSupiList,omitempty"`
	 Gpsi	string	`json:"gpsi,omitempty"`
	 GroupId	string	`json:"groupId,omitempty"`
	 ExcludeGpsiList	[]string	`json:"excludeGpsiList,omitempty"`
	 SubscribingNfType	NFType	`json:"subscribingNfType,omitempty"`
	 EventSyncInd	*bool	`json:"eventSyncInd,omitempty"`
	 BindingInfo	[]string	`json:"bindingInfo,omitempty"`
	 Pei	string	`json:"pei,omitempty"`
	 EventList	[]AmfEvent	`json:"eventList"`
	 NotifyCorrelationId	string	`json:"notifyCorrelationId"`
	 AoiStateList	map[string]AreaOfInterestEventState	`json:"aoiStateList,omitempty"`
}
