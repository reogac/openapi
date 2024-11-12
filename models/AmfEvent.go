package models

type AmfEvent struct {
	Type                   AmfEventType            `json:"type"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
}
