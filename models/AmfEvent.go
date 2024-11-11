package models

type AmfEvent struct {
	TrafficDescriptorList  []TrafficDescriptor `json:"trafficDescriptorList,omitempty"`
	ReportUeReachable      *bool               `json:"reportUeReachable,omitempty"`
	NextReport             string              `json:"nextReport,omitempty"`
	NextPeriodicReportTime string              `json:"nextPeriodicReportTime,omitempty"`
	ReachabilityFilter     string              `json:"reachabilityFilter,omitempty"`
	MaxReports             *int                `json:"maxReports,omitempty"`
	SnssaiFilter           []ExtSnssai         `json:"snssaiFilter,omitempty"`
	IdleStatusInd          *bool               `json:"idleStatusInd,omitempty"`
	LocationFilterList     []string            `json:"locationFilterList,omitempty"`
	UeInAreaFilter         *UeInAreaFilter     `json:"ueInAreaFilter,omitempty"`
	DispersionArea         *DispersionArea     `json:"dispersionArea,omitempty"`
	TargetArea             *TargetArea         `json:"targetArea,omitempty"`
	Type                   string              `json:"type"`
	ImmediateFlag          *bool               `json:"immediateFlag,omitempty"`
	AreaList               []AmfEventArea      `json:"areaList,omitempty"`
	RefId                  *int                `json:"refId,omitempty"`
	UdmDetectInd           *bool               `json:"udmDetectInd,omitempty"`
	PresenceInfoList       *presenceInfoList   `json:"presenceInfoList,omitempty"`
	MaxResponseTime        *int                `json:"maxResponseTime,omitempty"`
	MinInterval            *int                `json:"minInterval,omitempty"`
}
