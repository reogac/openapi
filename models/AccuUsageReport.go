package models

type AccuUsageReport struct {
	NextVolUsageUplink   *int64 `json:"nextVolUsageUplink,omitempty"`
	VolUsage             *int64 `json:"volUsage,omitempty"`
	VolUsageUplink       *int64 `json:"volUsageUplink,omitempty"`
	VolUsageDownlink     *int64 `json:"volUsageDownlink,omitempty"`
	TimeUsage            *int   `json:"timeUsage,omitempty"`
	NextVolUsage         *int64 `json:"nextVolUsage,omitempty"`
	NextVolUsageDownlink *int64 `json:"nextVolUsageDownlink,omitempty"`
	NextTimeUsage        *int   `json:"nextTimeUsage,omitempty"`
	RefUmIds             string `json:"refUmIds"`
}
