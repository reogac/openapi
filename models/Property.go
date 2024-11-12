package models

type Property struct {
	Regex    string `json:"regex,omitempty"`
	Value    string `json:"value,omitempty"`
	Name     string `json:"name"`
	Required *bool  `json:"required,omitempty"`
}
