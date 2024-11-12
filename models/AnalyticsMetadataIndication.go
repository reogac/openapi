package models

type AnalyticsMetadataIndication struct {
	DataWindow    *TimeWindow    `json:"dataWindow,omitempty"`
	DataStatProps []string       `json:"dataStatProps,omitempty"`
	Strategy      OutputStrategy `json:"strategy,omitempty"`
	AggrNwdafIds  []string       `json:"aggrNwdafIds,omitempty"`
}
