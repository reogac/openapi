package models

type ClassCriterion struct {
	DisperClass    string `json:"disperClass"`
	ClassThreshold int    `json:"classThreshold"`
	ThresMatch     string `json:"thresMatch"`
}
