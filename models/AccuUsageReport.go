package models
type AccuUsageReport struct {
	 NextVolUsage	*int64	`json:"nextVolUsage,omitempty"`
	 NextVolUsageDownlink	*int64	`json:"nextVolUsageDownlink,omitempty"`
	 RefUmIds	string	`json:"refUmIds"`
	 VolUsage	*int64	`json:"volUsage,omitempty"`
	 TimeUsage	*int	`json:"timeUsage,omitempty"`
	 NextVolUsageUplink	*int64	`json:"nextVolUsageUplink,omitempty"`
	 NextTimeUsage	*int	`json:"nextTimeUsage,omitempty"`
	 VolUsageUplink	*int64	`json:"volUsageUplink,omitempty"`
	 VolUsageDownlink	*int64	`json:"volUsageDownlink,omitempty"`
}
