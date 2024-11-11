package models

type DispersionCollection struct {
	Snssai         *Snssai             `json:"snssai,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	DisperClass    string              `json:"disperClass,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
	Supis          []string            `json:"supis,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	UsageRank      *int                `json:"usageRank,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
}
