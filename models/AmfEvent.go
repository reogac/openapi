package models

type AmfEvent struct {
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	Type                   AmfEventType            `json:"type"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
}
