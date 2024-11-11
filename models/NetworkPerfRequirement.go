package models

type NetworkPerfRequirement struct {
	RelativeRatio *int   `json:"relativeRatio,omitempty"`
	AbsoluteNum   *int   `json:"absoluteNum,omitempty"`
	NwPerfType    string `json:"nwPerfType"`
}
