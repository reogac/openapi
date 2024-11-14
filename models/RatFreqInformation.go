package models
type RatFreqInformation struct {
	 SvcExpThreshold	*ThresholdLevel	`json:"svcExpThreshold,omitempty"`
	 MatchingDir	MatchingDirection	`json:"matchingDir,omitempty"`
	 AllFreq	*bool	`json:"allFreq,omitempty"`
	 AllRat	*bool	`json:"allRat,omitempty"`
	 Freq	*int	`json:"freq,omitempty"`
	 RatType	RatType	`json:"ratType,omitempty"`
}
