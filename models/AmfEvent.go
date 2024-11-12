package models

type AmfEvent struct {
	Type                   AmfEventType            `json:"type"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
}
