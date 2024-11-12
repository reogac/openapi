package models

type NetworkPerfInfo struct {
	AbsoluteNum   *int             `json:"absoluteNum,omitempty"`
	Confidence    *int             `json:"confidence,omitempty"`
	NetworkArea   *NetworkAreaInfo `json:"networkArea,omitempty"`
	NwPerfType    NetworkPerfType  `json:"nwPerfType,omitempty"`
	RelativeRatio *int             `json:"relativeRatio,omitempty"`
}
