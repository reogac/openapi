package models

type AmfEvent struct {
	LocationFilterList    []string            `json:"locationFilterList,omitempty"`
	ReportUeReachable     *bool               `json:"reportUeReachable,omitempty"`
	ReachabilityFilter    string              `json:"reachabilityFilter,omitempty"`
	MaxReports            *int                `json:"maxReports,omitempty"`
	IdleStatusInd         *bool               `json:"idleStatusInd,omitempty"`
	AreaList              []AmfEventArea      `json:"areaList,omitempty"`
	ImmediateFlag         *bool               `json:"immediateFlag,omitempty"`
	RefId                 *int                `json:"refId,omitempty"`
	TrafficDescriptorList []TrafficDescriptor `json:"trafficDescriptorList,omitempty"`
	MaxResponseTime       *int                `json:"maxResponseTime,omitempty"`
	Type                  string              `json:"type"`
}
