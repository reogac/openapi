package models
type LocationInfo struct {
	 Loc	UserLocation	`json:"loc"`
	 Ratio	*int	`json:"ratio,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
}
