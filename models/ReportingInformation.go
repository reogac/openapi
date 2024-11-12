package models

type ReportingInformation struct {
	RepPeriod         *int               `json:"repPeriod,omitempty"`
	SampRatio         *int               `json:"sampRatio,omitempty"`
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
	MonDur            string             `json:"monDur,omitempty"`
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
	ImmRep            *bool              `json:"immRep,omitempty"`
}
