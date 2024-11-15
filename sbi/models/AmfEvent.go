/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	Type                   AmfEventType            `json:"type"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
}
