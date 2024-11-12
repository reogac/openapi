package models

type AccuUsageReport struct {
	RefUmIds             string `json:"refUmIds"`
	VolUsageUplink       *int64 `json:"volUsageUplink,omitempty"`
	TimeUsage            *int   `json:"timeUsage,omitempty"`
	NextVolUsageUplink   *int64 `json:"nextVolUsageUplink,omitempty"`
	NextTimeUsage        *int   `json:"nextTimeUsage,omitempty"`
	VolUsage             *int64 `json:"volUsage,omitempty"`
	VolUsageDownlink     *int64 `json:"volUsageDownlink,omitempty"`
	NextVolUsage         *int64 `json:"nextVolUsage,omitempty"`
	NextVolUsageDownlink *int64 `json:"nextVolUsageDownlink,omitempty"`
}
