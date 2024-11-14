/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	Type                   AmfEventType            `json:"type"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
}
