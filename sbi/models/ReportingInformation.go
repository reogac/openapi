/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReportingInformation struct {
	RepPeriod         *int               `json:"repPeriod,omitempty"`
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	SampRatio         *int               `json:"sampRatio,omitempty"`
	ImmRep            *bool              `json:"immRep,omitempty"`
	MonDur            string             `json:"monDur,omitempty"`
}
