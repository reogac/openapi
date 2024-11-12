package models

type ExtSnssai struct {
	SdRanges   []SdRange   `json:"sdRanges,omitempty"`
	WildcardSd *WildcardSd `json:"wildcardSd,omitempty"`
	Sst        int         `json:"sst"`
	Sd         string      `json:"sd,omitempty"`
}
