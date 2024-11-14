package models
type AmfEvent struct {
	 AreaList	[]AmfEventArea	`json:"areaList,omitempty"`
	 TargetArea	*TargetArea	`json:"targetArea,omitempty"`
	 UeInAreaFilter	*UeInAreaFilter	`json:"ueInAreaFilter,omitempty"`
	 DispersionArea	*DispersionArea	`json:"dispersionArea,omitempty"`
	 NextPeriodicReportTime	string	`json:"nextPeriodicReportTime,omitempty"`
	 Type	AmfEventType	`json:"type"`
	 ReachabilityFilter	ReachabilityFilter	`json:"reachabilityFilter,omitempty"`
	 NextReport	string	`json:"nextReport,omitempty"`
	 LocationFilterList	[]string	`json:"locationFilterList,omitempty"`
	 TrafficDescriptorList	[]TrafficDescriptor	`json:"trafficDescriptorList,omitempty"`
	 ReportUeReachable	*bool	`json:"reportUeReachable,omitempty"`
	 MaxResponseTime	*int	`json:"maxResponseTime,omitempty"`
	 MinInterval	*int	`json:"minInterval,omitempty"`
	 IdleStatusInd	*bool	`json:"idleStatusInd,omitempty"`
	 ImmediateFlag	*bool	`json:"immediateFlag,omitempty"`
	 RefId	*int	`json:"refId,omitempty"`
	 UdmDetectInd	*bool	`json:"udmDetectInd,omitempty"`
	 MaxReports	*int	`json:"maxReports,omitempty"`
	 PresenceInfoList	map[string]PresenceInfo	`json:"presenceInfoList,omitempty"`
	 SnssaiFilter	[]ExtSnssai	`json:"snssaiFilter,omitempty"`
}
