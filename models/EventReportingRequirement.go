package models

type EventReportingRequirement struct {
	Accuracy          Accuracy                     `json:"accuracy,omitempty"`
	EndTs             string                       `json:"endTs,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
	StartTs           string                       `json:"startTs,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
}
