package models

type AmfEvent struct {
	MaxReports             *int                    `json:"maxReports,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	Type                   AmfEventType            `json:"type"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
}
