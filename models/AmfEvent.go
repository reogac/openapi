package models

type AmfEvent struct {
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	Type                   AmfEventType            `json:"type"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
}
