package models

type AmfEvent struct {
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	Type                   AmfEventType            `json:"type"`
}
