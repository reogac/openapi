package models
type ExtSnssai struct {
	 WildcardSd	*WildcardSd	`json:"wildcardSd,omitempty"`
	 Sst	int	`json:"sst"`
	 Sd	string	`json:"sd,omitempty"`
	 SdRanges	[]SdRange	`json:"sdRanges,omitempty"`
}
