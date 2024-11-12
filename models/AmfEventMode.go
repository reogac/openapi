package models

type AmfEventMode struct {
	PartitioningCriteria []string         `json:"partitioningCriteria,omitempty"`
	NotifFlag            NotificationFlag `json:"notifFlag,omitempty"`
	Trigger              AmfEventTrigger  `json:"trigger"`
	MaxReports           *int             `json:"maxReports,omitempty"`
	Expiry               string           `json:"expiry,omitempty"`
	RepPeriod            *int             `json:"repPeriod,omitempty"`
	SampRatio            *int             `json:"sampRatio,omitempty"`
}
