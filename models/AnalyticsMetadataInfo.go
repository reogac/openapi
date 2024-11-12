package models

type AnalyticsMetadataInfo struct {
	DataStatProps []string    `json:"dataStatProps,omitempty"`
	Strategy      string      `json:"strategy,omitempty"`
	Accuracy      string      `json:"accuracy,omitempty"`
	NumSamples    *int        `json:"numSamples,omitempty"`
	DataWindow    *TimeWindow `json:"dataWindow,omitempty"`
}
