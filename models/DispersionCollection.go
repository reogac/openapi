package models
type DispersionCollection struct {
	 UeLoc	*UserLocation	`json:"ueLoc,omitempty"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 Supis	[]string	`json:"supis,omitempty"`
	 DisperAmount	*int	`json:"disperAmount,omitempty"`
	 UeRatio	*int	`json:"ueRatio,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 Gpsis	[]string	`json:"gpsis,omitempty"`
	 AppVolumes	[]ApplicationVolume	`json:"appVolumes,omitempty"`
	 DisperClass	DispersionClass	`json:"disperClass,omitempty"`
	 UsageRank	*int	`json:"usageRank,omitempty"`
	 PercentileRank	*int	`json:"percentileRank,omitempty"`
}
