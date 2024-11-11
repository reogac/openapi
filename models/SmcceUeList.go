package models

type SmcceUeList struct {
	MediumLevel []string `json:"mediumLevel,omitempty"`
	LowLevel    []string `json:"lowLevel,omitempty"`
	HighLevel   []string `json:"highLevel,omitempty"`
}
