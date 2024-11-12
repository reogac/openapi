package models

type ReportingInformation struct {
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	MonDur            string             `json:"monDur,omitempty"`
	SampRatio         *int               `json:"sampRatio,omitempty"`
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
	ImmRep            *bool              `json:"immRep,omitempty"`
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	RepPeriod         *int               `json:"repPeriod,omitempty"`
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
}
