package models
type AmfEvent struct {
	 PresenceInfoList	map[string]PresenceInfo	`json:"presenceInfoList,omitempty"`
	 IdleStatusInd	*bool	`json:"idleStatusInd,omitempty"`
	 Type	AmfEventType	`json:"type"`
	 RefId	*int	`json:"refId,omitempty"`
	 ReportUeReachable	*bool	`json:"reportUeReachable,omitempty"`
	 ReachabilityFilter	ReachabilityFilter	`json:"reachabilityFilter,omitempty"`
	 TrafficDescriptorList	[]TrafficDescriptor	`json:"trafficDescriptorList,omitempty"`
	 UdmDetectInd	*bool	`json:"udmDetectInd,omitempty"`
	 SnssaiFilter	[]ExtSnssai	`json:"snssaiFilter,omitempty"`
	 NextPeriodicReportTime	string	`json:"nextPeriodicReportTime,omitempty"`
	 MaxResponseTime	*int	`json:"maxResponseTime,omitempty"`
	 TargetArea	*TargetArea	`json:"targetArea,omitempty"`
	 UeInAreaFilter	*UeInAreaFilter	`json:"ueInAreaFilter,omitempty"`
	 NextReport	string	`json:"nextReport,omitempty"`
	 MinInterval	*int	`json:"minInterval,omitempty"`
	 DispersionArea	*DispersionArea	`json:"dispersionArea,omitempty"`
	 ImmediateFlag	*bool	`json:"immediateFlag,omitempty"`
	 AreaList	[]AmfEventArea	`json:"areaList,omitempty"`
	 LocationFilterList	[]string	`json:"locationFilterList,omitempty"`
	 MaxReports	*int	`json:"maxReports,omitempty"`
}
