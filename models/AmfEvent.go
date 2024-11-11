package models

type AmfEvent struct {
	MaxReports             *int                    `json:"maxReports,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	Type                   string                  `json:"type"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	ReachabilityFilter     string                  `json:"reachabilityFilter,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
}
