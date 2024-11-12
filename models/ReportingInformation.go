package models

type ReportingInformation struct {
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	RepPeriod         *int               `json:"repPeriod,omitempty"`
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
	ImmRep            *bool              `json:"immRep,omitempty"`
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	MonDur            string             `json:"monDur,omitempty"`
	SampRatio         *int               `json:"sampRatio,omitempty"`
}
