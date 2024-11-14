/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	Type                   AmfEventType            `json:"type"`
	RefId                  *int                    `json:"refId,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
}
