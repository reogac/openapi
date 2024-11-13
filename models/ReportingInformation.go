package models

type ReportingInformation struct {
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
	SampRatio         *int               `json:"sampRatio,omitempty"`
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
	ImmRep            *bool              `json:"immRep,omitempty"`
	MonDur            string             `json:"monDur,omitempty"`
	RepPeriod         *int               `json:"repPeriod,omitempty"`
}
