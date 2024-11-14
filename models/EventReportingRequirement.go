/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventReportingRequirement struct {
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	StartTs           string                       `json:"startTs,omitempty"`
	EndTs             string                       `json:"endTs,omitempty"`
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
	Accuracy          Accuracy                     `json:"accuracy,omitempty"`
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
}
