package models

type DispersionCollection struct {
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
	Snssai         *Snssai             `json:"snssai,omitempty"`
	Supis          []string            `json:"supis,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	UsageRank      *int                `json:"usageRank,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	DisperClass    DispersionClass     `json:"disperClass,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
}
