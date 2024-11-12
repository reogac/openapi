package models

type AccuUsageReport struct {
	VolUsageDownlink     *int64 `json:"volUsageDownlink,omitempty"`
	NextVolUsage         *int64 `json:"nextVolUsage,omitempty"`
	NextVolUsageUplink   *int64 `json:"nextVolUsageUplink,omitempty"`
	NextTimeUsage        *int   `json:"nextTimeUsage,omitempty"`
	RefUmIds             string `json:"refUmIds"`
	VolUsage             *int64 `json:"volUsage,omitempty"`
	VolUsageUplink       *int64 `json:"volUsageUplink,omitempty"`
	TimeUsage            *int   `json:"timeUsage,omitempty"`
	NextVolUsageDownlink *int64 `json:"nextVolUsageDownlink,omitempty"`
}
