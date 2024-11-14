package models
type AmfEventSubscription struct {
	 NotifyCorrelationId	string	`json:"notifyCorrelationId"`
	 NfId	string	`json:"nfId"`
	 SubsChangeNotifyUri	string	`json:"subsChangeNotifyUri,omitempty"`
	 SubsChangeNotifyCorrelationId	string	`json:"subsChangeNotifyCorrelationId,omitempty"`
	 Supi	string	`json:"supi,omitempty"`
	 IncludeSupiList	[]string	`json:"includeSupiList,omitempty"`
	 Gpsi	string	`json:"gpsi,omitempty"`
	 EventNotifyUri	string	`json:"eventNotifyUri"`
	 SourceNfType	NFType	`json:"sourceNfType,omitempty"`
	 ExcludeGpsiList	[]string	`json:"excludeGpsiList,omitempty"`
	 IncludeGpsiList	[]string	`json:"includeGpsiList,omitempty"`
	 ExcludeSupiList	[]string	`json:"excludeSupiList,omitempty"`
	 Pei	string	`json:"pei,omitempty"`
	 Options	*AmfEventMode	`json:"options,omitempty"`
	 GroupId	string	`json:"groupId,omitempty"`
	 AnyUE	*bool	`json:"anyUE,omitempty"`
	 EventList	[]AmfEvent	`json:"eventList"`
}
