package models

type EventReportingRequirement struct {
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	EndTs             string                       `json:"endTs,omitempty"`
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
	StartTs           string                       `json:"startTs,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	Accuracy          Accuracy                     `json:"accuracy,omitempty"`
}
