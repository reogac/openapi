/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccuUsageReport struct {
	TimeUsage            *int   `json:"timeUsage,omitempty"`
	NextVolUsageUplink   *int64 `json:"nextVolUsageUplink,omitempty"`
	NextTimeUsage        *int   `json:"nextTimeUsage,omitempty"`
	VolUsage             *int64 `json:"volUsage,omitempty"`
	VolUsageUplink       *int64 `json:"volUsageUplink,omitempty"`
	VolUsageDownlink     *int64 `json:"volUsageDownlink,omitempty"`
	NextVolUsage         *int64 `json:"nextVolUsage,omitempty"`
	NextVolUsageDownlink *int64 `json:"nextVolUsageDownlink,omitempty"`
	RefUmIds             string `json:"refUmIds"`
}
