/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionCollection struct {
	Supis          []string            `json:"supis,omitempty"`
	UsageRank      *int                `json:"usageRank,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	DisperClass    DispersionClass     `json:"disperClass,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
	Snssai         *Snssai             `json:"snssai,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
}
