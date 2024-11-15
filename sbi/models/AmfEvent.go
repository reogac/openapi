/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	Type                   AmfEventType            `json:"type"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
}
