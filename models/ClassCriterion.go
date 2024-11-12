package models

type ClassCriterion struct {
	ThresMatch     MatchingDirection `json:"thresMatch"`
	DisperClass    DispersionClass   `json:"disperClass"`
	ClassThreshold int               `json:"classThreshold"`
}
