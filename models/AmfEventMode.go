package models
type AmfEventMode struct {
	 Trigger	string	`json:"trigger"`
	 MaxReports	*int	`json:"maxReports,omitempty"`
	 Expiry	string	`json:"expiry,omitempty"`
	 RepPeriod	*int	`json:"repPeriod,omitempty"`
	 SampRatio	*int	`json:"sampRatio,omitempty"`
}
