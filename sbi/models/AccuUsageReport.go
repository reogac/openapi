/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccuUsageReport struct {
	NextTimeUsage        *int   `json:"nextTimeUsage,omitempty"`
	RefUmIds             string `json:"refUmIds"`
	VolUsage             *int64 `json:"volUsage,omitempty"`
	TimeUsage            *int   `json:"timeUsage,omitempty"`
	NextVolUsageUplink   *int64 `json:"nextVolUsageUplink,omitempty"`
	NextVolUsageDownlink *int64 `json:"nextVolUsageDownlink,omitempty"`
	VolUsageUplink       *int64 `json:"volUsageUplink,omitempty"`
	VolUsageDownlink     *int64 `json:"volUsageDownlink,omitempty"`
	NextVolUsage         *int64 `json:"nextVolUsage,omitempty"`
}
