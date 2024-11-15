/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	Type                   AmfEventType            `json:"type"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
}
