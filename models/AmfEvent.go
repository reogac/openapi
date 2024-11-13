package models

type AmfEvent struct {
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	Type                   AmfEventType            `json:"type"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
}
