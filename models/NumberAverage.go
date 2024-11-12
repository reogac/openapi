package models
type NumberAverage struct {
	 Number	float64	`json:"number"`
	 Variance	float64	`json:"variance"`
	 Skewness	*float64	`json:"skewness,omitempty"`
}
