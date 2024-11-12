package models

type ReportingInformation struct {
	ImmRep            *bool              `json:"immRep,omitempty"`
	RepPeriod         *int               `json:"repPeriod,omitempty"`
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	MonDur            string             `json:"monDur,omitempty"`
	SampRatio         *int               `json:"sampRatio,omitempty"`
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
}
