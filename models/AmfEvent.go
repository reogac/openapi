package models
type AmfEvent struct {
	 AreaList	[]AmfEventArea	`json:"areaList,omitempty"`
	 LocationFilterList	[]string	`json:"locationFilterList,omitempty"`
	 ReportUeReachable	*bool	`json:"reportUeReachable,omitempty"`
	 ImmediateFlag	*bool	`json:"immediateFlag,omitempty"`
	 RefId	*int	`json:"refId,omitempty"`
	 TrafficDescriptorList	[]TrafficDescriptor	`json:"trafficDescriptorList,omitempty"`
	 ReachabilityFilter	string	`json:"reachabilityFilter,omitempty"`
	 MaxReports	*int	`json:"maxReports,omitempty"`
	 MaxResponseTime	*int	`json:"maxResponseTime,omitempty"`
	 IdleStatusInd	*bool	`json:"idleStatusInd,omitempty"`
	 Type	string	`json:"type"`
}
