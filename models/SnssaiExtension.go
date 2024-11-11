package models

type SnssaiExtension struct {
	SdRanges   []SdRange `json:"sdRanges,omitempty"`
	WildcardSd *bool     `json:"wildcardSd,omitempty"`
}
