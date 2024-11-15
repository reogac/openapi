/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:12 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDeleteData struct {
	UeTimeZone           string             `json:"ueTimeZone,omitempty"`
	ServingNetwork       *PlmnIdNid         `json:"servingNetwork,omitempty"`
	UserLocationInfoTime string             `json:"userLocationInfoTime,omitempty"`
	RanNasRelCauses      []RanNasRelCause   `json:"ranNasRelCauses,omitempty"`
	AccuUsageReports     []AccuUsageReport  `json:"accuUsageReports,omitempty"`
	PduSessRelCause      PduSessionRelCause `json:"pduSessRelCause,omitempty"`
	UserLocationInfo     *UserLocation      `json:"userLocationInfo,omitempty"`
}
