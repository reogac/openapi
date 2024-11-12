package models
type SnssaiExtension struct {
	 SdRanges	[]SdRange	`json:"sdRanges,omitempty"`
	 WildcardSd	*WildcardSd	`json:"wildcardSd,omitempty"`
}
