package models

type AnalyticsMetadataIndication struct {
	DataStatProps []string       `json:"dataStatProps,omitempty"`
	Strategy      OutputStrategy `json:"strategy,omitempty"`
	AggrNwdafIds  []string       `json:"aggrNwdafIds,omitempty"`
	DataWindow    *TimeWindow    `json:"dataWindow,omitempty"`
}
