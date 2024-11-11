package models

type NetworkPerfInfo struct {
	RelativeRatio *int             `json:"relativeRatio,omitempty"`
	AbsoluteNum   *int             `json:"absoluteNum,omitempty"`
	Confidence    *int             `json:"confidence,omitempty"`
	NetworkArea   *NetworkAreaInfo `json:"networkArea,omitempty"`
	NwPerfType    string           `json:"nwPerfType,omitempty"`
}
