/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ThresholdLevel struct {
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
	CongLevel         *int     `json:"congLevel,omitempty"`
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
}
