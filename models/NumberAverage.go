package models

type NumberAverage struct {
	Skewness *float64 `json:"skewness,omitempty"`
	Number   float64  `json:"number"`
	Variance float64  `json:"variance"`
}
