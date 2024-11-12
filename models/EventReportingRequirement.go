package models

type EventReportingRequirement struct {
	Accuracy          Accuracy                     `json:"accuracy,omitempty"`
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
	StartTs           string                       `json:"startTs,omitempty"`
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	EndTs             string                       `json:"endTs,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
}
