package models

type NetworkPerfRequirement struct {
	NwPerfType    string `json:"nwPerfType"`
	RelativeRatio *int   `json:"relativeRatio,omitempty"`
	AbsoluteNum   *int   `json:"absoluteNum,omitempty"`
}
