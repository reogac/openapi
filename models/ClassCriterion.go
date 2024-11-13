package models

type ClassCriterion struct {
	DisperClass    DispersionClass   `json:"disperClass"`
	ClassThreshold int               `json:"classThreshold"`
	ThresMatch     MatchingDirection `json:"thresMatch"`
}
