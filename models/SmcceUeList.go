package models

type SmcceUeList struct {
	HighLevel   []string `json:"highLevel,omitempty"`
	MediumLevel []string `json:"mediumLevel,omitempty"`
	LowLevel    []string `json:"lowLevel,omitempty"`
}
