/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficInformation struct {
	UplinkRate     string `json:"uplinkRate,omitempty"`
	DownlinkRate   string `json:"downlinkRate,omitempty"`
	UplinkVolume   *int64 `json:"uplinkVolume,omitempty"`
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`
	TotalVolume    *int64 `json:"totalVolume,omitempty"`
}