package models

type DispersionCollection struct {
	UsageRank      *int                `json:"usageRank,omitempty"`
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
	Snssai         *Snssai             `json:"snssai,omitempty"`
	Supis          []string            `json:"supis,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	DisperClass    DispersionClass     `json:"disperClass,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
}
