package models
type NetworkPerfRequirement struct {
	 AbsoluteNum	*int	`json:"absoluteNum,omitempty"`
	 NwPerfType	NetworkPerfType	`json:"nwPerfType"`
	 RelativeRatio	*int	`json:"relativeRatio,omitempty"`
}
