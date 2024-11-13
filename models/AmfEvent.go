package models

type AmfEvent struct {
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	Type                   AmfEventType            `json:"type"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
}
