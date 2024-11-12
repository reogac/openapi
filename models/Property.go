package models

type Property struct {
	Name     string `json:"name"`
	Required *bool  `json:"required,omitempty"`
	Regex    string `json:"regex,omitempty"`
	Value    string `json:"value,omitempty"`
}
