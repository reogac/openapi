package models

type Property struct {
	Value    string `json:"value,omitempty"`
	Name     string `json:"name"`
	Required *bool  `json:"required,omitempty"`
	Regex    string `json:"regex,omitempty"`
}
