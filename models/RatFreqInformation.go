package models

type RatFreqInformation struct {
	AllRat          *bool           `json:"allRat,omitempty"`
	Freq            *int            `json:"freq,omitempty"`
	RatType         string          `json:"ratType,omitempty"`
	SvcExpThreshold *ThresholdLevel `json:"svcExpThreshold,omitempty"`
	MatchingDir     string          `json:"matchingDir,omitempty"`
	AllFreq         *bool           `json:"allFreq,omitempty"`
}
