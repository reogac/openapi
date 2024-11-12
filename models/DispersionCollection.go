package models

type DispersionCollection struct {
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	UsageRank      *int                `json:"usageRank,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
	Snssai         *Snssai             `json:"snssai,omitempty"`
	Supis          []string            `json:"supis,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	DisperClass    DispersionClass     `json:"disperClass,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
}
