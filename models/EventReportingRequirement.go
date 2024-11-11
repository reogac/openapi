package models

type EventReportingRequirement struct {
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
	StartTs           string                       `json:"startTs,omitempty"`
	EndTs             string                       `json:"endTs,omitempty"`
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
	Accuracy          string                       `json:"accuracy,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
}
