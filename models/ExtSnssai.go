package models

type ExtSnssai struct {
	Sd         string      `json:"sd,omitempty"`
	WildcardSd *WildcardSd `json:"wildcardSd,omitempty"`
	SdRanges   []SdRange   `json:"sdRanges,omitempty"`
	Sst        int         `json:"sst"`
}
