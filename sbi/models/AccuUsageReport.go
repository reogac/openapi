/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccuUsageReport struct {
	NextVolUsage         *int64 `json:"nextVolUsage,omitempty"`
	NextTimeUsage        *int   `json:"nextTimeUsage,omitempty"`
	VolUsage             *int64 `json:"volUsage,omitempty"`
	VolUsageUplink       *int64 `json:"volUsageUplink,omitempty"`
	VolUsageDownlink     *int64 `json:"volUsageDownlink,omitempty"`
	NextVolUsageDownlink *int64 `json:"nextVolUsageDownlink,omitempty"`
	RefUmIds             string `json:"refUmIds"`
	TimeUsage            *int   `json:"timeUsage,omitempty"`
	NextVolUsageUplink   *int64 `json:"nextVolUsageUplink,omitempty"`
}
