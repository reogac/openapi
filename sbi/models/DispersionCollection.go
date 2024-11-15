/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionCollection struct {
	DisperClass    DispersionClass     `json:"disperClass,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
	Snssai         *Snssai             `json:"snssai,omitempty"`
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	UsageRank      *int                `json:"usageRank,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
	Supis          []string            `json:"supis,omitempty"`
}
