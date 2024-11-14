package models
type ExtSnssai struct {
	 SdRanges	[]SdRange	`json:"sdRanges,omitempty"`
	 Sst	int	`json:"sst"`
	 Sd	string	`json:"sd,omitempty"`
	 WildcardSd	*WildcardSd	`json:"wildcardSd,omitempty"`
}
