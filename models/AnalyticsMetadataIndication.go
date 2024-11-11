package models

type AnalyticsMetadataIndication struct {
	Strategy      string      `json:"strategy,omitempty"`
	AggrNwdafIds  []string    `json:"aggrNwdafIds,omitempty"`
	DataWindow    *TimeWindow `json:"dataWindow,omitempty"`
	DataStatProps []string    `json:"dataStatProps,omitempty"`
}
