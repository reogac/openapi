package models

type AccuUsageReport struct {
	NextVolUsageUplink   *int64 `json:"nextVolUsageUplink,omitempty"`
	NextVolUsageDownlink *int64 `json:"nextVolUsageDownlink,omitempty"`
	NextTimeUsage        *int   `json:"nextTimeUsage,omitempty"`
	VolUsage             *int64 `json:"volUsage,omitempty"`
	TimeUsage            *int   `json:"timeUsage,omitempty"`
	VolUsageDownlink     *int64 `json:"volUsageDownlink,omitempty"`
	NextVolUsage         *int64 `json:"nextVolUsage,omitempty"`
	RefUmIds             string `json:"refUmIds"`
	VolUsageUplink       *int64 `json:"volUsageUplink,omitempty"`
}
