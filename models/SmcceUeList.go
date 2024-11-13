package models

type SmcceUeList struct {
	LowLevel    []string `json:"lowLevel,omitempty"`
	HighLevel   []string `json:"highLevel,omitempty"`
	MediumLevel []string `json:"mediumLevel,omitempty"`
}
