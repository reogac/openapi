/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReportingInformation struct {
	MonDur            string             `json:"monDur,omitempty"`
	RepPeriod         *int               `json:"repPeriod,omitempty"`
	SampRatio         *int               `json:"sampRatio,omitempty"`
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
	ImmRep            *bool              `json:"immRep,omitempty"`
}
