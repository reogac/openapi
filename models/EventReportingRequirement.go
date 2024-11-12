package models

type EventReportingRequirement struct {
	Accuracy          Accuracy                     `json:"accuracy,omitempty"`
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	StartTs           string                       `json:"startTs,omitempty"`
	EndTs             string                       `json:"endTs,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
}
