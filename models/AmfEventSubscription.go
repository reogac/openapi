package models
type AmfEventSubscription struct {
	 GroupId	string	`json:"groupId,omitempty"`
	 ExcludeGpsiList	[]string	`json:"excludeGpsiList,omitempty"`
	 IncludeSupiList	[]string	`json:"includeSupiList,omitempty"`
	 IncludeGpsiList	[]string	`json:"includeGpsiList,omitempty"`
	 Pei	string	`json:"pei,omitempty"`
	 AnyUE	*bool	`json:"anyUE,omitempty"`
	 EventList	[]AmfEvent	`json:"eventList"`
	 EventNotifyUri	string	`json:"eventNotifyUri"`
	 NotifyCorrelationId	string	`json:"notifyCorrelationId"`
	 SubsChangeNotifyUri	string	`json:"subsChangeNotifyUri,omitempty"`
	 SubsChangeNotifyCorrelationId	string	`json:"subsChangeNotifyCorrelationId,omitempty"`
	 Supi	string	`json:"supi,omitempty"`
	 SourceNfType	NFType	`json:"sourceNfType,omitempty"`
	 ExcludeSupiList	[]string	`json:"excludeSupiList,omitempty"`
	 Options	*AmfEventMode	`json:"options,omitempty"`
	 NfId	string	`json:"nfId"`
	 Gpsi	string	`json:"gpsi,omitempty"`
}
