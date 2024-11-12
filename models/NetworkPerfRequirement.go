package models

type NetworkPerfRequirement struct {
	NwPerfType    NetworkPerfType `json:"nwPerfType"`
	RelativeRatio *int            `json:"relativeRatio,omitempty"`
	AbsoluteNum   *int            `json:"absoluteNum,omitempty"`
}
