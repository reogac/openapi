package models

type ReportingInformation struct {
	NotifMethod       string   `json:"notifMethod,omitempty"`
	SampRatio         *int     `json:"sampRatio,omitempty"`
	NotifFlag         string   `json:"notifFlag,omitempty"`
	GrpRepTime        *int     `json:"grpRepTime,omitempty"`
	ImmRep            *bool    `json:"immRep,omitempty"`
	MaxReportNbr      *int     `json:"maxReportNbr,omitempty"`
	MonDur            string   `json:"monDur,omitempty"`
	RepPeriod         *int     `json:"repPeriod,omitempty"`
	PartitionCriteria []string `json:"partitionCriteria,omitempty"`
}
