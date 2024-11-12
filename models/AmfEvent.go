package models

type AmfEvent struct {
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	ReachabilityFilter     string                  `json:"reachabilityFilter,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	Type                   string                  `json:"type"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
}
