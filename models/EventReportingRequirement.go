package models
type EventReportingRequirement struct {
	 OffsetPeriod	*int	`json:"offsetPeriod,omitempty"`
	 MaxObjectNbr	*int	`json:"maxObjectNbr,omitempty"`
	 MaxSupiNbr	*int	`json:"maxSupiNbr,omitempty"`
	 AnaMeta	[]string	`json:"anaMeta,omitempty"`
	 HistAnaTimePeriod	*TimeWindow	`json:"histAnaTimePeriod,omitempty"`
	 Accuracy	Accuracy	`json:"accuracy,omitempty"`
	 StartTs	string	`json:"startTs,omitempty"`
	 EndTs	string	`json:"endTs,omitempty"`
	 SampRatio	*int	`json:"sampRatio,omitempty"`
	 TimeAnaNeeded	string	`json:"timeAnaNeeded,omitempty"`
	 AnaMetaInd	*AnalyticsMetadataIndication	`json:"anaMetaInd,omitempty"`
	 AccPerSubset	[]string	`json:"accPerSubset,omitempty"`
}
