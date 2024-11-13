package models

type ReportingInformation struct {
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	MonDur            string             `json:"monDur,omitempty"`
	RepPeriod         *int               `json:"repPeriod,omitempty"`
	SampRatio         *int               `json:"sampRatio,omitempty"`
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
	ImmRep            *bool              `json:"immRep,omitempty"`
}
