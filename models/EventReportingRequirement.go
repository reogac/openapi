package models

type EventReportingRequirement struct {
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
	EndTs             string                       `json:"endTs,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	Accuracy          Accuracy                     `json:"accuracy,omitempty"`
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
	StartTs           string                       `json:"startTs,omitempty"`
}
