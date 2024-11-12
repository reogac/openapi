package models

type SvcExperience struct {
	LowerRange *float64 `json:"lowerRange,omitempty"`
	Mos        *float64 `json:"mos,omitempty"`
	UpperRange *float64 `json:"upperRange,omitempty"`
}
