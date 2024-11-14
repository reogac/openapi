package models
type DispersionCollection struct {
	 Gpsis	[]string	`json:"gpsis,omitempty"`
	 DisperAmount	*int	`json:"disperAmount,omitempty"`
	 DisperClass	DispersionClass	`json:"disperClass,omitempty"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 Supis	[]string	`json:"supis,omitempty"`
	 UsageRank	*int	`json:"usageRank,omitempty"`
	 PercentileRank	*int	`json:"percentileRank,omitempty"`
	 UeRatio	*int	`json:"ueRatio,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 UeLoc	*UserLocation	`json:"ueLoc,omitempty"`
	 AppVolumes	[]ApplicationVolume	`json:"appVolumes,omitempty"`
}
