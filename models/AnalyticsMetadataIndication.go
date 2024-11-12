package models

type AnalyticsMetadataIndication struct {
	AggrNwdafIds  []string       `json:"aggrNwdafIds,omitempty"`
	DataWindow    *TimeWindow    `json:"dataWindow,omitempty"`
	DataStatProps []string       `json:"dataStatProps,omitempty"`
	Strategy      OutputStrategy `json:"strategy,omitempty"`
}
