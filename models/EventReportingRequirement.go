package models
type EventReportingRequirement struct {
	 AccPerSubset	[]string	`json:"accPerSubset,omitempty"`
	 EndTs	string	`json:"endTs,omitempty"`
	 MaxObjectNbr	*int	`json:"maxObjectNbr,omitempty"`
	 AnaMetaInd	*AnalyticsMetadataIndication	`json:"anaMetaInd,omitempty"`
	 HistAnaTimePeriod	*TimeWindow	`json:"histAnaTimePeriod,omitempty"`
	 Accuracy	Accuracy	`json:"accuracy,omitempty"`
	 StartTs	string	`json:"startTs,omitempty"`
	 OffsetPeriod	*int	`json:"offsetPeriod,omitempty"`
	 SampRatio	*int	`json:"sampRatio,omitempty"`
	 MaxSupiNbr	*int	`json:"maxSupiNbr,omitempty"`
	 TimeAnaNeeded	string	`json:"timeAnaNeeded,omitempty"`
	 AnaMeta	[]string	`json:"anaMeta,omitempty"`
}
