/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficInformation struct {
	DownlinkRate   string `json:"downlinkRate,omitempty"`
	UplinkVolume   *int64 `json:"uplinkVolume,omitempty"`
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`
	TotalVolume    *int64 `json:"totalVolume,omitempty"`
	UplinkRate     string `json:"uplinkRate,omitempty"`
}
