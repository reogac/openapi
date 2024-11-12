package models
type AmfEvent struct {
	 LocationFilterList	[]string	`json:"locationFilterList,omitempty"`
	 RefId	*int	`json:"refId,omitempty"`
	 UdmDetectInd	*bool	`json:"udmDetectInd,omitempty"`
	 MaxReports	*int	`json:"maxReports,omitempty"`
	 UeInAreaFilter	*UeInAreaFilter	`json:"ueInAreaFilter,omitempty"`
	 IdleStatusInd	*bool	`json:"idleStatusInd,omitempty"`
	 AreaList	[]AmfEventArea	`json:"areaList,omitempty"`
	 TrafficDescriptorList	[]TrafficDescriptor	`json:"trafficDescriptorList,omitempty"`
	 ReachabilityFilter	ReachabilityFilter	`json:"reachabilityFilter,omitempty"`
	 NextReport	string	`json:"nextReport,omitempty"`
	 ImmediateFlag	*bool	`json:"immediateFlag,omitempty"`
	 MaxResponseTime	*int	`json:"maxResponseTime,omitempty"`
	 MinInterval	*int	`json:"minInterval,omitempty"`
	 NextPeriodicReportTime	string	`json:"nextPeriodicReportTime,omitempty"`
	 Type	AmfEventType	`json:"type"`
	 ReportUeReachable	*bool	`json:"reportUeReachable,omitempty"`
	 PresenceInfoList	map[string]PresenceInfo	`json:"presenceInfoList,omitempty"`
	 TargetArea	*TargetArea	`json:"targetArea,omitempty"`
	 SnssaiFilter	[]ExtSnssai	`json:"snssaiFilter,omitempty"`
	 DispersionArea	*DispersionArea	`json:"dispersionArea,omitempty"`
}
