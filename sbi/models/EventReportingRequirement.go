/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventReportingRequirement struct {
	Accuracy          Accuracy                     `json:"accuracy,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
	StartTs           string                       `json:"startTs,omitempty"`
	EndTs             string                       `json:"endTs,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
}
