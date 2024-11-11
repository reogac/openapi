package models

type ExtSnssai struct {
	Sd         string    `json:"sd,omitempty"`
	Sst        int       `json:"sst"`
	SdRanges   []SdRange `json:"sdRanges,omitempty"`
	WildcardSd *bool     `json:"wildcardSd,omitempty"`
}
