package models
type AccuUsageReport struct {
	 VolUsageUplink	*int64	`json:"volUsageUplink,omitempty"`
	 NextVolUsage	*int64	`json:"nextVolUsage,omitempty"`
	 NextVolUsageDownlink	*int64	`json:"nextVolUsageDownlink,omitempty"`
	 VolUsage	*int64	`json:"volUsage,omitempty"`
	 VolUsageDownlink	*int64	`json:"volUsageDownlink,omitempty"`
	 TimeUsage	*int	`json:"timeUsage,omitempty"`
	 NextVolUsageUplink	*int64	`json:"nextVolUsageUplink,omitempty"`
	 NextTimeUsage	*int	`json:"nextTimeUsage,omitempty"`
	 RefUmIds	string	`json:"refUmIds"`
}
