package models

type NumberAverage struct {
	Variance float64  `json:"variance"`
	Skewness *float64 `json:"skewness,omitempty"`
	Number   float64  `json:"number"`
}
