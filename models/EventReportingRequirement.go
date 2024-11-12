package models

type EventReportingRequirement struct {
	StartTs           string                       `json:"startTs,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
	Accuracy          string                       `json:"accuracy,omitempty"`
	EndTs             string                       `json:"endTs,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
}
