package models

type DispersionCollection struct {
	UsageRank      *int                `json:"usageRank,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
	Snssai         *Snssai             `json:"snssai,omitempty"`
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	DisperClass    string              `json:"disperClass,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
	Supis          []string            `json:"supis,omitempty"`
}
