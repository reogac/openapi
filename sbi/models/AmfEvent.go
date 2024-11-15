/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:28 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	Type                   AmfEventType            `json:"type"`
	RefId                  *int                    `json:"refId,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
}
