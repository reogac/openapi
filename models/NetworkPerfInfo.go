package models

type NetworkPerfInfo struct {
	Confidence    *int             `json:"confidence,omitempty"`
	NetworkArea   *NetworkAreaInfo `json:"networkArea,omitempty"`
	NwPerfType    string           `json:"nwPerfType,omitempty"`
	RelativeRatio *int             `json:"relativeRatio,omitempty"`
	AbsoluteNum   *int             `json:"absoluteNum,omitempty"`
}
