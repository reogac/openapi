/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PerfData struct {
	MaxPacketDelay    *int   `json:"maxPacketDelay,omitempty"`
	AvgPacketLossRate *int   `json:"avgPacketLossRate,omitempty"`
	AvgTrafficRate    string `json:"avgTrafficRate,omitempty"`
	MaxTrafficRate    string `json:"maxTrafficRate,omitempty"`
	AvePacketDelay    *int   `json:"avePacketDelay,omitempty"`
}
