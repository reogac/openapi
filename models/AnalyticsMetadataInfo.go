package models
type AnalyticsMetadataInfo struct {
	 DataWindow	*TimeWindow	`json:"dataWindow,omitempty"`
	 DataStatProps	[]string	`json:"dataStatProps,omitempty"`
	 Strategy	OutputStrategy	`json:"strategy,omitempty"`
	 Accuracy	Accuracy	`json:"accuracy,omitempty"`
	 NumSamples	*int	`json:"numSamples,omitempty"`
}
