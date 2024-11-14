/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionCollection struct {
	UsageRank      *int                `json:"usageRank,omitempty"`
	Snssai         *Snssai             `json:"snssai,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	DisperClass    DispersionClass     `json:"disperClass,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
	Supis          []string            `json:"supis,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
}
