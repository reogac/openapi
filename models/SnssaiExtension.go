package models

type SnssaiExtension struct {
	WildcardSd *WildcardSd `json:"wildcardSd,omitempty"`
	SdRanges   []SdRange   `json:"sdRanges,omitempty"`
}
